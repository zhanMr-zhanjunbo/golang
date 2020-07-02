package router

import (
	"github.com/gin-gonic/gin"
	"ginTest/controller"
)

func SetUpRouter() *gin.Engine{
	r :=gin.Default()
	r.Static("/static","./static")
	r.LoadHTMLGlob("./templates/*")
	r.GET("/", controller.GetIndex)

	r.GET("/billList", controller.GetBillList)
	billGroup:=r.Group("/billList")
	{
		billGroup.GET("/billAdd", controller.GetBillAdd)
		billGroup.GET("/billUpdate/:orderId",controller.GetBillUpdate )
		billGroup.POST("/billAdd",controller.PostBillAdd)
		billGroup.GET("/billView/:orderId",controller.GetBillOneView)
		billGroup.PUT("/billUpdate/:orderId",controller.UpdateBillOneView)
		billGroup.DELETE("/billView/:orderId",controller.DeleteBillOneView)
		billGroup.GET("/billFindMore",controller.GetMoreBill)
	}

	r.GET("/providerList", controller.GetProviderList)
	proGroup:=r.Group("/providerList")
	{
		proGroup.GET("/providerView", controller.GetProvideView)
		proGroup.GET("/providerAdd", controller.GetProvideAdd)
		proGroup.GET("/providerUpdate", controller.GetProvideUpdate)
	}

	r.GET("/userList", controller.GetUserList)
	userGroup:=r.Group("/userList")
	{
		userGroup.GET("/userView", controller.GetUserView)
		userGroup.GET("/userAdd", controller.GetUserAdd)
		userGroup.GET("/userUpdate", controller.GetUserUpdate)
	}

	r.GET("/password",controller.GetPassword )
	r.GET("/login", controller.GetLogin)
	return r
}
