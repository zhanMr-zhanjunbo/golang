package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"golang_normal_study/websocket/impl"
	"net/http"
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
	if wsConn, err = upGrader.Upgrade(w, r, nil); err != nil {
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
		fmt.Print("read from client is: ", string(buf))
		if err = conn.WriteMessage(buf); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()

}
func main() {
	fmt.Println("server listen is start:")
	http.HandleFunc("/echo", GetWebSocket)
	if err := http.ListenAndServe("0.0.0.0:9000", nil); err != nil {
		fmt.Println("监听端口9000失败:", err.Error())
	}
}
