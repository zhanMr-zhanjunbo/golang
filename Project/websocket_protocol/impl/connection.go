package impl

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
	"encoding/json"
	"fmt"
)

var (
	wsUsers map[*websocket.Conn]string=make(map[*websocket.Conn]string)
)
type Connection struct {
	wsConn       *websocket.Conn
	//wsUsers      map[*websocket.Conn]string
	outChannel   chan []byte
	inChannel    chan []byte
	closeChannel chan []byte
	mutex        sync.Mutex
	isClosed     bool
}

func InitCreateConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:       wsConn,
		//wsUsers:      make(map[*websocket.Conn]string,1024),
		outChannel:   make(chan []byte, 1024),
		inChannel:    make(chan []byte, 1024),
		closeChannel: make(chan []byte, 1),
	}
	wsUsers[wsConn]=""
	go conn.ReadLoop()
	go conn.WriteLoop()
	return
}
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChannel:
	case <-conn.closeChannel:
		err = errors.New("connection is closed")
	}
	return
}
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChannel <- data:
	case <-conn.closeChannel:
		err = errors.New("connection is closed")
	}
	return
}
func (conn *Connection) Close() {
	conn.wsConn.Close()
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChannel)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}
func (conn *Connection) ReadLoop() {
	var data []byte
	var err error
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			delete(wsUsers,conn.wsConn)
			goto ERR
		}
		conn.ParseData(data)
		select {
		case conn.inChannel <- data:
		case <-conn.closeChannel:
			goto ERR
		}
	}
ERR:
	conn.Close()
}
func (conn *Connection)ParseData(data []byte) {
	var (
		tidings Tidings
		dataNews DataSet
		err error
	)
	dataNews=DataSet{}
	if err=json.Unmarshal(data,&tidings);err!=nil{
		fmt.Println("parse data failed",err.Error())
		return
	}
	if _,ok:=wsUsers[conn.wsConn];ok{
		wsUsers[conn.wsConn]+=tidings.Name
		dataNews.Users=append(dataNews.Users,User{Name:tidings.Name})
	}
	dataNews.News=append(dataNews.News,tidings)
}
func (conn *Connection) WriteLoop() {
	var data []byte
	var err error
	for {
		select {
		case data = <-conn.outChannel:
		case <-conn.closeChannel:
			goto ERR
		}
		for wsConn:=range wsUsers{
			if err = wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
				delete(wsUsers,conn.wsConn)
				goto ERR
			}
		}
	}
ERR:
	conn.Close()
}
