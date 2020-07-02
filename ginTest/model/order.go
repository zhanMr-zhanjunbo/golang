package model

import (
	"ginTest/dao"
	"github.com/pkg/errors"
	"time"
)

//订单：Order
//Id
//订单编码    OrderCode
//商品名称    GoodsName
//商品单位    GoodsUnit
//商品数量    GoodsNumbers
//总金额      TotalAmount
//供应商      SupplierName
//支付状态    PayStatus  0:未支付 1：已支付
type Orders struct {
	Id int               `json:"_" `
	OrderCode string     `json:"order_code" binding:"required"`
	GoodsName string     `json:"goods_name" binding:"required"`
	GoodsUnit string     `json:"goods_unit" binding:"required"`
	GoodsNumbers int     `json:"goods_numbers" binding:"required"`
	TotalAmount  int     `json:"total_amount" binding:"required"`
	SupplierName string  `json:"supplier_name" binding:"required"`
	PayStatus    int     `json:"pay_status" binding:"required"`
	CreateTime   string  `json:"_"`
}
func FindSameOrderCode(orderCode string)(n int64){
	var order []Orders
	n=dao.DB.Where("",orderCode).Find(&order).RowsAffected
	return
}
func AddOrders(order Orders) (err error){
	  order.CreateTime=time.Now().Format("2006-01-02 15:04:05")
      if err=dao.DB.Create(&order).Error;err!=nil{
      	err=errors.New("add to mysql failed")
      	return
	  }
	  return
}
func FindAllBill() (order []Orders,err error) {
	if err=dao.DB.Find(&order).Error;err!=nil{
		return nil,errors.New("not find")
	}
	return order,nil
}
func FindOneBill(id string)(order Orders,err error){
	if err=dao.DB.Debug().Where("order_code=?",id).Find(&order).Error;err!=nil{
		err=errors.New("find failed")
		return
	}
	return
}
func UpdateOneBill(id string)(order Orders,err error){
	if err=dao.DB.Debug().Where("order_code=?",id).First(&order).Error;err!=nil{
		err=errors.New("modify failed")
		return
	}
	return
}
func SaveOneBill(order Orders) (err error) {
	if err=dao.DB.Debug().Save(order).Error;err!=nil{
		err=errors.New("save failed")
		return
	}
	return
}
func DeleteOneBill(id string) (err error)  {
	var order Orders
	if err=dao.DB.Where("order_code=?",id).Delete(&order).Error;err!=nil{
		err=errors.New("delete failed")
		return
	}
	return
}
func FindMoreBill(goodsName ,supplier string,payStatus int)(order []Orders,err error){
	db:=dao.DB
	if len(goodsName)!=0{
		db=db.Where("goods_name=?",goodsName)
	}
	if len(supplier)!=0{
		db=db.Where("supplier_name=?",supplier)
	}
	if payStatus!=0{
		db=db.Where("pay_status=?",payStatus)
	}
	if err=db.Debug().Find(&order).Error;err!=nil{
		return nil,errors.New("find failed")
	}
	return order,nil
}