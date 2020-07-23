package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"Project/websocket_protocol/impl"
	"net/http"
	"html/template"
	"log"
	"time"
)

var (
	upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func GetWebSocket(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		conn   *impl.Connection
		err    error
		buf    []byte
	)
	if r.Header.Get("origin")!="http://"+r.Host{
		http.Error(w,"origin not allowed",403)
		return
	}
	if wsConn, err = upGrader.Upgrade(w, r, w.Header()); err != nil {
		return
	}
	if conn, err = impl.InitCreateConnection(wsConn); err != nil {
		goto ERR
	}
	go func() {
		for {
			var err error
			if err = conn.WriteMessage([]byte("heartbeat\n")); err != nil {
				return
			}
			time.Sleep(2 * time.Second)
		}

	}()
	for {
		if buf, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		//fmt.Println("read from client is: ", string(buf))
		if err = conn.WriteMessage(buf); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()

}
func GetIndex(w http.ResponseWriter,r *http.Request){
	t, _ := template.ParseFiles("./views/index2.html")
	log.Println(r.Host)
	t.Execute(w,"http://"+r.Host+"/echo")
}
func main() {
	fmt.Println("server listen is start:")
	http.HandleFunc("/echo", GetWebSocket)
	http.HandleFunc("/",GetIndex)
	if err := http.ListenAndServe("0.0.0.0:9000", nil); err != nil {
		fmt.Println("监听端口9000失败:", err.Error())
	}
}
