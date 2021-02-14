package session

//定义管理者,管理所有的session
type SessionMgr interface {
	// 初始化
	Init(addr string , options ...string)(err error)
	CreateSession()(seesion Session, err error)
	Get(sessionId string)(seesion Session, err error)
}
