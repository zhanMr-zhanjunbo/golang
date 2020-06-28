package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ginSession/model"
	"ginSession/dao"
)
func AccessIndex(c *gin.Context)  {
		c.HTML(http.StatusOK,"index.html",nil)
}
func Register(c *gin.Context) {
	var user model.UserRegister
	c.BindJSON(&user)
	if !model.CheckIsRegister(user.UserName){
		c.JSON(http.StatusOK,gin.H{"msg":"用户已注册"})
	}else {
		if err:=model.DB.Create(&user).Error;err!=nil{
			c.JSON(http.StatusOK,gin.H{"msg":err.Error()})
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
		c.JSON(http.StatusOK,gin.H{"msg":err.Error()})
	}else{
		tmpSD,ok:=c.Get("session")
		if !ok{
			panic("session middleware")
		}
		sd:=tmpSD.(dao.SessionData)
		sd.Set("isLogin",true)
		sd.Save()
		c.JSON(http.StatusOK,gin.H{
			"msg":"login success",
			"data":user,
		})
	}
}
func Home(c *gin.Context)  {
	c.HTML(http.StatusOK,"home.html",nil)
}