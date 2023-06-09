package api

import "github.com/gin-gonic/gin"

func initImageFamiliesRoutes(rg *gin.RouterGroup) {
	rg.GET("", getImageFamilies)
	rg.POST("", createImageFamily)
	rg.GET("/:id", getImageFamily)
	rg.PUT("/:id", updateImageFamily)
	rg.DELETE("/:id", deleteImageFamily)
}

func getImageFamilies(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func createImageFamily(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func getImageFamily(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func updateImageFamily(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func deleteImageFamily(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}
