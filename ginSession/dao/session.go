package dao
type SessionData interface {
	GetID()(sessionId string)
	Get(key string)(value interface{}, err error)
	Set(key string, value interface{})(err error)
	Del(key string) (err error)
	Save()(err error)
}
