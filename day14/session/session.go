package session

// session中间件

// 定义一个Session的接口
type Session interface {
	Set(key string , value interface{}) error
	Get(key string ) (interface{}, error)
	Del(key string ) error
	Save() error
}