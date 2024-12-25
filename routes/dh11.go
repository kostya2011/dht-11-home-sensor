package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTemperature(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Dummy!",
	})
}

func getHumidity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Dummy!",
	})
}
