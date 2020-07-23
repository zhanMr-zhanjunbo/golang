package main

import (
	"net/http"
	"golang.org/x/net/websocket"
	"log"
	"strings"
	"html/template"
	"os"
)
func upper(ws *websocket.Conn){
	var err error
	for{
		var reply string
		if err = websocket.Message.Receive(ws, &reply);err!=nil{
			log.Printf("receive is error,%s\n",err.Error())
			continue
		}
		if err = websocket.Message.Send(ws, strings.ToUpper(reply));err!=nil{
			log.Printf("send is error,%s\n",err.Error())
			continue
		}
	}

}
func index(w http.ResponseWriter,r *http.Request)  {
	method := r.Method
	log.Println("请求类型是：", method)
	t, _ := template.ParseFiles("./views/index.html")
	t.Execute(w,nil)
}
func main(){     
    http.Handle("/upper",websocket.Handler(upper))
	http.HandleFunc("/",index)
	log.Println("正在监听9999端口.....")
    if err:=http.ListenAndServe(":9999",nil);err!=nil{
    	log.Printf("listen is error,%s\n",err.Error())
    	os.Exit(1)
	}
}

