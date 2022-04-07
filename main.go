package main

import (
	"OwnWeb/tool"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	InitialSetting(app)
}

func InitialSetting(app *gin.Engine) {
	config, err := tool.ReadConfigFile("application")
	if err != nil {
		panic("读取配置失败")
	}
	tool.RegistRouter(app, "Router")
	app.Run(":" + config.App.Port)
}
