package main

import (
	"bufio"
	"fmt"
	"golang_normal_study/go_log/loginit"
	"net"
	"os"
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
func MessageSend(conn net.Conn) {
	fmt.Println("client send message start:")
	buf := make([]byte, 1024)
	var n int
	var err error
	for {
		reader := bufio.NewReader(os.Stdin)
		if n, err = reader.Read(buf); err != nil {
			conn.Close()
			break
		}
		conn.Write(buf[:n])
	}
}
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	CheckErr(err)
	defer conn.Close()
	go MessageSend(conn)
	for {
		reader := bufio.NewReader(conn)
		buf := make([]byte, 1024)
		var n int
		var err error
		for {
			if n, err = reader.Read(buf); err != nil {
				conn.Close()
				break
			}
			fmt.Print("read from server is:", string(buf[:n]))
		}
	}
}
