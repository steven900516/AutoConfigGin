package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloController struct {
}

func (helloController *HelloController) Router(app *gin.Engine) {
	app.GET("/hello", helloController.hello)
}

func (HelloController *HelloController) hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"path":    "/hello",
		"message": "success router",
	})
}
