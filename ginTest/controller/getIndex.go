package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK,"index.html",nil)
}
