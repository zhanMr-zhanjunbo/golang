package dao

import (
	"sync"
	"github.com/astaxie/goredis"
	"strconv"
	"encoding/json"
	"fmt"
	"errors"
	"github.com/satori/go.uuid"
)
type RedisManager struct {
	Session map[string]SessionData
	lock sync.RWMutex
	client goredis.Client
}
func (r *RedisManager)Init(addr string,option ...string)  {
	var (
		password string
		db string
		dbValue int
		err error
	)
	if len(option)==1{
		password=option[0]
	}else if len(option)==2{
		password=option[0]
		db=option[1]
	}
	if dbValue,err=strconv.Atoi(db);err!=nil{
		dbValue=0
	}
	r.client.Addr=addr
	r.client.Password=password
	r.client.Db=dbValue
	if _,err=r.client.Ping();err!=nil{
		panic(err)
	}
}
func (r *RedisManager)LoadFromRedis(sessionId string) (err error) {
	var value []byte
	if value,err=r.client.Get(sessionId);err!=nil{
		fmt.Println("from redis get failed",err.Error())
		return
	}
	if err=json.Unmarshal(value,&r.Session);err!=nil{
		fmt.Println("json analysis failed",err.Error())
		return
	}
	return
}
func (r *RedisManager)GetSessionData(sessionId string)(sd SessionData,err error) {
	if r.Session==nil{
		if err=r.LoadFromRedis(sessionId);err!=nil{
			return
		}
	}
	r.lock.RLock()
	defer r.lock.RUnlock()
	sd,ok:=r.Session[sessionId]
	if !ok{
		err=errors.New("invalid sessionId")
		return
	}
	return
}
func (r *RedisManager)CreateSession()(sd SessionData,err error) {
	uid:=uuid.NewV4()
	sessionId:=uid.String()
	sd=NewRedisSession(sessionId,r.client)
	r.Session[sd.GetID()]=sd
	return
}
func NewRedisManager() Manager{
	return &RedisManager{
		Session:make(map[string]SessionData,1024),
	}
}