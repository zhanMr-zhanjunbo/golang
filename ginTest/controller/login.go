package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK,"login.html",nil)
}
func GetPassword(c *gin.Context) {
	c.HTML(http.StatusOK,"password.html",nil)
}