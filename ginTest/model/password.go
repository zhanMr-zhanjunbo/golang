package model
//密码：Password
//旧密码    OldPassword
//新密码    NewPassword
//确认密码  ConfirmPassword
type Password struct {
	Id int
	OldPassword     string
	NewPassword     string
	ConfirmPassword	string
}
