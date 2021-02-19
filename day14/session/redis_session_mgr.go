package session

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

// RedisSessionMgr设计
// 定义RedisSessionMgr对象(字段:存放所有的session 的map, 读写锁)
// 构造函数
//Init()
// CreateSession()
// GetSession()

type RedisSessionMgr struct {
	// redis 地址
	addr string
	// 密码
	passwd string
	// 连接池
	pool *redis.Pool
	// 锁
	rwlock sync.RWMutex
	// 大map
	sessionMap map[string]Session
}

// 构造
func NewRedisSeesionMgr() *RedisSessionMgr {
	sr := &RedisSessionMgr{
		sessionMap: make(map[string]Session, 32),
	}
	return sr
}

// 创建连接池
func myPool(addr, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   64,
		MaxActive: 1000,
		IdleTimeout: 240*time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			// 若有密码,判断
			if _, err := conn.Do("AUTH", password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		// 连接测试,开发时写
		// 上线注释掉
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

// 初始化
func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	// 若有其他参数
	if len(options) >0{
		r.passwd = options[0]
	}
	// 创建连接池
	r.pool = myPool(addr, r.passwd)
	r.addr = addr
	return err
}

// 创建session
func (r *RedisSessionMgr) CreateSession()(session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	//使用uuid   go get github.com/satori/go.uuid
	// 用uuid 作为sessionId
	id := uuid.NewV4()
	// 转成string
	sessionId := id.String()
	// 创建个session
	session = NewRedisSession(sessionId, r.pool)
	// 把单个session 加入到大map中
	r.sessionMap[sessionId] = session
	return
}

// 获取session
func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}
