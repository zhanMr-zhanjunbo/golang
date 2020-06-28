package impl

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	wsConn       *websocket.Conn
	outChannel   chan []byte
	inChannel    chan []byte
	closeChannel chan []byte
	mutex        sync.Mutex
	isClosed     bool
}

func InitCreateConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:       wsConn,
		outChannel:   make(chan []byte, 1024),
		inChannel:    make(chan []byte, 1024),
		closeChannel: make(chan []byte, 1),
	}
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
			goto ERR
		}
		select {
		case conn.inChannel <- data:
		case <-conn.closeChannel:
			goto ERR
		}
	}
ERR:
	conn.Close()
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
		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
