package loginit

import (
	"log"
	"os"
)

var Logger *log.Logger
var LogFile *os.File

func init() {
	var err error
	LogFile, err = os.OpenFile("./testlog", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	Logger = log.New(LogFile, "\r\n", log.Ldate|log.Ltime)
	//Logger.Println("hello worlddddddd!")   //Println 是普通打印
	//Logger.Panic("")   //Panic 是打印panic,可以通过recover()捕获
	//Logger.Fatal("")   //Fatal 是终止退出程序,相当于os.Exit(1)异常退出
}
