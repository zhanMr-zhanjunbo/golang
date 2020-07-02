package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context) {
	c.HTML(http.StatusOK,"userList.html",nil)
}
func GetUserView(c *gin.Context) {
	c.HTML(http.StatusOK,"userView.html",nil)
}
func GetUserAdd(c *gin.Context) {
	c.HTML(http.StatusOK,"userAdd.html",nil)
}
func GetUserUpdate(c *gin.Context) {
	c.HTML(http.StatusOK,"userUpdate.html",nil)
}