package dao

type Manager interface {
	Init(addr string,option...string)
	CreateSession()(sd SessionData,err error)
	GetSessionData(sessionId string)(sd SessionData,err error)
}
