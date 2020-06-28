package main

import "golang_normal_study/go_log/loginit"

/*
var logger *log.Logger
var logFile *os.File
func init() {
     var err error
	logFile,err=os.OpenFile("./testlog.log",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err!=nil{
		panic(err)
	}
	//defer logFile.Close()
	logger=log.New(logFile,"\r\n",log.Ldate|log.Ltime)
	//writer=bufio.NewWriter(logFile)
	//log.SetOutput(writer)
	//log.Println("hello121222212112")
	//log.Println("helloerrerrrrr")
	//writer.Flush()
}
*/
func Add(x, y int) {
	sum := x + y
	loginit.Logger.Println("两数之和：", sum)
}
func Product(x, y int) {
	product := x * y
	loginit.Logger.Println("两数乘积：", product)
}
func main() {
	defer loginit.LogFile.Close() //init初始化不能关闭文件，否则写不进去
	loginit.Logger.Println("111111111")
	Add(33, 10)
	loginit.Logger.Println("222222222")
	Product(2, 3)
	loginit.Logger.Println("3333333333")
}
