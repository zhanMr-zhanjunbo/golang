package main

import (
	"github.com/gin-gonic/gin"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"ginSession/controller"
	_"ginSession/model"
	"ginSession/model"
	"ginSession/proto"
	"ginSession/dao"
)

func main() {
	defer model.DB.Close()
	r :=gin.Default()
	r.Static("/static","./static")
	r.LoadHTMLGlob("./templates/*")
	dao.InitMgr("redis","127.0.0.1:6379")
	r.Use(proto.IsLogin(dao.MgrObj))
	r.GET("/index",controller.AccessIndex)
	r.POST("/index/register",controller.Register)
	r.POST("/index/login", controller.Login)
	r.GET("/home", proto.AuthMiddleware,controller.Home)
	r.Run(":9090")

}
