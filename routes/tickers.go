package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTickerValue(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Dummy!",
	})
}

func getRegisteredTickers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Dummy!",
	})
}

func registerTicker(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Dummy!",
	})
}

func unregisterTicker(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Dummy!",
	})
}
