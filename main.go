package main

import (
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
	_ = engine.Run(":" + global.App.Configuration.App.Port)
}
