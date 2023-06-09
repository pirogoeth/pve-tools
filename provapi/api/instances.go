package api

import "github.com/gin-gonic/gin"

func initInstancesRoutes(rg *gin.RouterGroup) {
	rg.GET("", getInstances)
	rg.POST("", createInstance)
	rg.GET("/:id", getInstance)
	rg.PUT("/:id", updateInstance)
	rg.DELETE("/:id", deleteInstance)
}

func getInstances(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func createInstance(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func getInstance(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func updateInstance(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func deleteInstance(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}
