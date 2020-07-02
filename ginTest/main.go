package main

import (
	"ginTest/router"
	"ginTest/dao"
	"ginTest/model"
)

func main() {
	if err := dao.InitMysql();err!=nil{
		panic(err)
	}
	defer dao.Close()
	dao.DB.SingularTable(true)
	dao.DB.AutoMigrate(&model.Orders{})
	r:=router.SetUpRouter()
	if err:=r.Run(":9090");err!=nil{
		panic(err)
	}
}
