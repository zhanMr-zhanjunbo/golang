package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
)

var (
	dialer = websocket.Dialer{}
	url    = "ws://127.0.0.1:9000/echo"
	conn   *websocket.Conn
	err    error
)

func WriteToServer(conn *websocket.Conn) {
	fmt.Println("client write message to server:")
	bytes := make([]byte, 1024)
	reader := bufio.NewReader(os.Stdin)
	for {
		var n int
		if n, err = reader.Read(bytes); err != nil {
			conn.Close()
			break
		}
		if err = conn.WriteMessage(websocket.TextMessage, []byte(bytes[:n])); err != nil {
			conn.Close()
			break
		}
	}
}
func main() {
	var buff []byte
	if conn, _, err = dialer.Dial(url, nil); err != nil {
		goto ERR
	}
	go WriteToServer(conn)
	for {
		if _, buff, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		fmt.Print("read from server is:", string(buff))
	}
ERR:
	conn.Close()
}
