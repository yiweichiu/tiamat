package routes

import (
	"net/http"
	"tiamat/m/v0/services"

	"github.com/gin-gonic/gin"
)

func Entry() {
	services.Init()

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})
	r.POST("/callback", services.Handler)

	r.Run()
}
