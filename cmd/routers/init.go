package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Test Endpoint")
	})
}
