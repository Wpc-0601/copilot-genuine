package main

import (
	"fmt"
	"github.com/Wpc-0601/copilot-genuine/config/bootstrap"
	"github.com/Wpc-0601/copilot-genuine/config/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	bootstrap.InitializeConfig()
	engine := gin.Default()
	engine.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "This is an new Application...")
	})
	fmt.Println("port is :", global.App.Configuration.App.Port)
	fmt.Println("env is: ", global.App.Configuration.App.Env)
	fmt.Println("appName is: ", global.App.Configuration.App.AppName)
	fmt.Println("appUrl is: ", global.App.Configuration.App.AppUrl)
	_ = engine.Run(":" + global.App.Configuration.App.Port)
}
