package session

// 2021-02-14 20:19:45
// 有点乱 老师说明天讲
import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

// 使用uuid   go get github.com/satori/go.uuid
// 定义一个memory_session_mgr对象
type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock sync.RWMutex
}

// 构造函数
func NewMemorySessionMgr ()*MemorySessionMgr{
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}

// 初始化
func (s *MemorySessionMgr) Init(addr string, options ...string)(err error){
	return err
}

// 创建
func (s *MemorySessionMgr) CreateSession()(session Session , err error){
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	//使用uuid   go get github.com/satori/go.uuid
	// 用uuid 作为sessionId
	id := uuid.NewV4()
	// 转成string
	sessionId := id.String()
	// 创建个session
	session = NewMemorySession(sessionId)
	// 把单个session 加入到大map中
	s.sessionMap[sessionId] = session
	return
}

func  (s *MemorySessionMgr)Get(sessionId string)(session Session, err error)  {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	session, ok := s.sessionMap[sessionId]
	if !ok{
		err = errors.New("session not exists")
		return
	}
	return
}

