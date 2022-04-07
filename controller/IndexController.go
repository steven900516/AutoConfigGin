package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (indexController *IndexController) Router(app *gin.Engine) {
	app.GET("/index", indexController.index)
}

func (indexController *IndexController) index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"path":    "/index",
		"message": "success router",
	})
}
