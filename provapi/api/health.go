package api

import (
	"github.com/gin-gonic/gin"
)

func initHealthRoutes(router *gin.RouterGroup) {
	router.GET("", getHealth)
}

func getHealth(ctx *gin.Context) {
	// todo: check health of clients too
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}
