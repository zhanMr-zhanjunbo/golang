package dao

import (
	"sync"
	"github.com/astaxie/goredis"
	"github.com/pkg/errors"
	"encoding/json"
)

type RedisSession struct {
	SessionId string
	SessionMap map[string]interface{}
	lock sync.RWMutex
	expired int64
	client goredis.Client
}
func NewRedisSession(sessionId string,client goredis.Client)*RedisSession{
	return &RedisSession{
		SessionId:sessionId,
		SessionMap:make(map[string]interface{},1024),
		expired:3600,
		client:client,
	}
}
func (r *RedisSession)GetID()(sessionId string) {
	sessionId=r.SessionId
	return
}
func (r *RedisSession)Get(key string)(value interface{},err error){
	r.lock.RLock()
	defer r.lock.RUnlock()
	value,ok:=r.SessionMap[key]
	if !ok{
		err=errors.New("invalid key!")
		return
	}
	return
}
func (r *RedisSession)Set(key string,value interface{})(err error)  {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.SessionMap[key]=value
	return
}
func (r *RedisSession)Del(key string)(err error)  {
	r.lock.Lock()
	defer r.lock.Unlock()
	delete(r.SessionMap,key)
	return
}
func (r *RedisSession)Save() (err error) {
	var value []byte
	if value,err=json.Marshal(r.SessionMap);err!=nil{
		err=errors.New("redis 序列化failed!")
		return
	}
	r.client.Setex(r.SessionId,r.expired,value)
	return
}