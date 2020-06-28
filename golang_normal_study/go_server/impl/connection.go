package impl

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

type Connection struct {
	tcpConn      net.Conn
	conns        *map[string]net.Conn
	inChannel    chan []byte
	outChannel   chan []byte
	closeChannel chan []byte
	mutex        sync.Mutex
	isClosed     bool
}

func InitCreateConnection(tcpConn net.Conn, tcpConns *map[string]net.Conn) (conn *Connection, err error) {
	conn = &Connection{
		tcpConn:      tcpConn,
		conns:        tcpConns,
		inChannel:    make(chan []byte, 1024),
		outChannel:   make(chan []byte, 1024),
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
	conn.tcpConn.Close()
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChannel)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}
func (conn *Connection) ReadLoop() {
	var n int
	var err error
	buf := make([]byte, 1024)
	for {
		if n, err = conn.tcpConn.Read(buf); err != nil {
			delete(*conn.conns, conn.tcpConn.RemoteAddr().String())
			goto ERR
		}
		select {
		case conn.inChannel <- buf[:n]:
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
	var key string
	var connObject net.Conn
	for {
		select {
		case data = <-conn.outChannel:
		case <-conn.closeChannel:
			goto ERR
		}
		for key, connObject = range *conn.conns {
			fmt.Println("connection connected is:", key, connObject)
			if _, err = connObject.Write(data); err != nil {
				delete(*conn.conns, key)
				goto ERR
			}
		}
	}
ERR:
	conn.Close()
}
