package proto

import (
	"github.com/gin-gonic/gin"
	"ginSession/dao"
	"net/http"
	"fmt"
)

func IsLogin(mgr dao.Manager) gin.HandlerFunc  {
	return func(c *gin.Context) {
		var sd dao.SessionData
		id,err:=c.Cookie("session_id")
		if err!=nil{
			sd,_=mgr.CreateSession()
		}else {
			sd,err=mgr.GetSessionData(id)
			if err!=nil{
				sd,_=mgr.CreateSession()
				id=sd.GetID()
			}
		}
		c.Set("session",sd)
		c.SetCookie("session_id",id,3600,"/","localhost",false,true)
		c.Next()
	}
}
func AuthMiddleware(c *gin.Context){
	// 1. 从上下文中取到session data
	// 1. 先从上下文中获取session data
	fmt.Println("in Auth")
	tmpSD, _ := c.Get("session")
	sd := tmpSD.(dao.SessionData)
	// 2. 从session data取到isLogin
	fmt.Printf("%#v\n", sd)
	value, err := sd.Get("isLogin")
	if err != nil {
		fmt.Println(err)
		// 取不到就是没有登录
		c.Redirect(http.StatusFound, "/index")
		return
	}
	fmt.Println(value)
	isLogin, ok := value.(bool)//类型断言
	if !ok {
		fmt.Println("!ok")
		c.Redirect(http.StatusFound, "/index")
		return
	}
	fmt.Println(isLogin)
	if !isLogin{
		c.Redirect(http.StatusFound, "/index")
		return
	}
	c.Next()
}

