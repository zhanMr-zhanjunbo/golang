package main

import (
	"fmt"
	"golang_normal_study/go_log/loginit"
	"golang_normal_study/go_server/impl"
	"net"
	"time"
)

func CheckErr(err error) {
	defer func() {
		if wrong, ok := recover().(error); ok {
			loginit.Logger.Println("程序出现异常：", wrong.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}
func Handler(tcpConn net.Conn, conns *map[string]net.Conn) {
	var conn *impl.Connection
	var err error
	var data []byte
	if conn, err = impl.InitCreateConnection(tcpConn, conns); err != nil {
		goto ERR
	}
	go func() {
		var err error
		for {
			if err = conn.WriteMessage([]byte("heartbeat\n")); err != nil {
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()
	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}

func main() {

	defer loginit.LogFile.Close()
	ln, err1 := net.Listen("tcp", "127.0.0.1:9000")
	CheckErr(err1)
	loginit.Logger.Println("监听开始：")
	defer ln.Close()
	conns := make(map[string]net.Conn)
	for {
		tcpConn, err := ln.Accept()
		if err != nil {
			loginit.Logger.Println("发生侦听错误：", err.Error())
			continue
		}
		fmt.Println("client connect address is:", tcpConn.RemoteAddr().String())
		conns[tcpConn.RemoteAddr().String()] = tcpConn
		go Handler(tcpConn, &conns)
	}
}
