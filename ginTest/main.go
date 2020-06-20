package main

import (
	"github.com/gin-gonic/gin"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"ginTest/controller"
	_"ginTest/model"
	"ginTest/model"
)

func main() {
	defer model.DB.Close()
	r :=gin.Default()
	r.Static("/static","./static")
	r.LoadHTMLGlob("./templates/*")
	r.GET("/index",controller.AccessIndex)
	r.POST("/index/register",controller.Register)
	r.POST("/index/login", controller.Login)
	r.Run(":9090")
}
