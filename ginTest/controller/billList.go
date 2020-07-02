package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ginTest/model"
	"strconv"
)

func GetBillList(c *gin.Context) {
	order, err := model.FindAllBill()
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"msg":"not find",
		})
	}
	c.HTML(http.StatusOK,"billList.html",gin.H{
		"data":order,
	})
}
func GetBillAdd(c *gin.Context) {
	c.HTML(http.StatusOK,"billAdd.html",nil)
}
func GetBillUpdate(c *gin.Context) {
	orderId:=c.Param("orderId")
	order, err := model.FindOneBill(orderId)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"msg":"find failed",
		})
	}
	c.HTML(http.StatusOK,"billUpdate.html",gin.H{
		"data":order,
	})
}
func GetMoreBill(c *gin.Context)  {
	goodsName:=c.Query("goods_name")
	supplier:=c.Query("supplier_name")
	payStatus:=c.Query("pay_status")
	py,_:=strconv.Atoi(payStatus)
	order, err := model.FindMoreBill(goodsName, supplier, py)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"not find",
		})
	}
	c.String(http.StatusOK,"data",order)
}
func PostBillAdd(c *gin.Context) {
	var (
		order model.Orders
		err error
	)
	if err=c.ShouldBind(&order);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"json invalid failed",
		})
	}
	if n:=model.FindSameOrderCode(order.OrderCode);n!=0{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"not sane orderCode",
		})
		return
	}
	if err=model.AddOrders(order);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"add failed!",
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"msg":"add success!",
		})
	}

}
func GetBillOneView(c *gin.Context)  {
	orderId:=c.Param("orderId")
	order, err := model.FindOneBill(orderId)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"msg":"find failed",
		})
	}
	c.HTML(http.StatusOK,"billView.html",gin.H{
		"data":order,
	})
}
func UpdateBillOneView(c *gin.Context){
	orderId:=c.Param("orderId")
	order, err := model.UpdateOneBill(orderId)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error":"modify failed",
		})
	}else{
		c.ShouldBindJSON(&order)
		if err:=model.SaveOneBill(order);err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"error":"save failed",
			})
		}
		c.JSON(http.StatusOK,gin.H{
			"msg":"modify success",
		})
	}
}
func DeleteBillOneView(c *gin.Context){
	orderId:=c.Param("orderId")
	if err := model.DeleteOneBill(orderId);err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"msg":"delete failed",
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"msg":"delete success",
		})
	}
}