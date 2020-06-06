package models

import "github.com/astaxie/goredis"

const(
	URL_QUEUE="url_queue"
	URL_VISIT_QUEUE="url_visit_queue"
)
var (
	client goredis.Client
)

func CreateConnect(addr string) {
	client.Addr=addr
}
func GetQueueLength()int{
	n,err:=client.Llen(URL_QUEUE)
	if err!=nil{
		return 0
	}
	return n
}
func PushInQueue(url string){
	client.Lpush(URL_QUEUE,[]byte(url))
}
func PopOutQueue() string{
	buf,err:=client.Rpop(URL_QUEUE)
	if err!=nil{
		panic(err)
	}
	return string(buf)
}
func AddToSet(url string) {
	client.Sadd(URL_VISIT_QUEUE,[]byte(url))
}
func IsVisit(url string)bool{
    isVisit,err:=client.Sismember(URL_VISIT_QUEUE,[]byte(url))
    if err!=nil{
    	return false
	}
	return isVisit
}

