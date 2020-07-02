package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProviderList(c *gin.Context) {
	c.HTML(http.StatusOK,"providerList.html",nil)
}
func GetProvideView(c *gin.Context) {
	c.HTML(http.StatusOK,"providerView.html",nil)
}
func GetProvideAdd(c *gin.Context) {
	c.HTML(http.StatusOK,"providerAdd.html",nil)
}
func GetProvideUpdate(c *gin.Context) {
	c.HTML(http.StatusOK,"providerUpdate.html",nil)
}