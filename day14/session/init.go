package session

import "fmt"

// 中间件
// 定义规范
// 现有设计需求文档,,,然后再开发写代码
// 文档看README里面

var (
	sessionMgr SessionMgr
)


func Init(provider string, addr string, options ...string)(err error){
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSeesionMgr()
	default:
		fmt.Errorf("不支持")
		return err
	}
	err = sessionMgr.Init(addr, options...)
	return err
}