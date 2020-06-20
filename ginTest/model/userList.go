package model

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)
type UserRegister struct {
	Id int          `gorm:"primary_key;auto_increment" json:"id" form:"id"`
	UserName string `json:"username" form:"username"`
	Email string    `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
func init(){
	var err error
	DB,err=gorm.Open("mysql","root:microsys@tcp(localhost:3306)/test?charset=utf8mb4")
	if err!=nil{
		panic(err)
	}
	DB.SingularTable(true)
	DB.AutoMigrate(&UserRegister{})
}
func CheckIsRegister(username string) bool{
	var user UserRegister
	if err:=DB.Where("user_name=?",username).Find(&user).Error;err==nil{
		return false
	}
	return true
}