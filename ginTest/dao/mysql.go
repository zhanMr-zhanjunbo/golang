package dao

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	)

var(
	DB *gorm.DB
)

func InitMysql() (err error)  {
	DB,err=gorm.Open("mysql","root:microsys@tcp(localhost:3306)/test?charset=utf8mb4")
	if err!=nil{
		return
	}
	return DB.DB().Ping()
}
func Close()  {
	DB.Close()
}