package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ginTest/model"
)
func AccessIndex(c *gin.Context)  {
		c.HTML(http.StatusOK,"index.html",nil)
}
func Register(c *gin.Context) {
	var user model.UserRegister
	c.BindJSON(&user)
	if !model.CheckIsRegister(user.UserName){
		c.JSON(http.StatusOK,gin.H{"error":"用户已注册"})
	}else {
		if err:=model.DB.Create(&user).Error;err!=nil{
			c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"msg":"register success",
			})
		}
	}
}
func Login(c *gin.Context) {
	var user model.UserRegister
	c.BindJSON(&user)
	if err:=model.DB.Where("email=? and password=?",user.Email,user.Password).Find(&user).Error;err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"msg":"login success",
			"data":user,
		})
	}
}