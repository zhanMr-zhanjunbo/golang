package model
//用户信息：User
//用户编码  UserCode
//用户名称  UserName
//用户密码  Password
//确认密码  ConfirmPassword
//用户性别  Sex
//出生日期  Birthday
//用户电话  PhoneNumber
//用户地址  Address
//用户类型  UserType 0:管理员  1:经理  2:普通用户
type User struct {
	id int
	UserCode string
	UserName string
	Password string
	ConfirmPassword string
	Sex             string
	Birthday        string
	PhoneNumber     int
	Address         string
	UserType	    int
}

