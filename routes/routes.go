package routes

import (
	"github.com/gin-gonic/gin"
)

func NewRoutes(server *gin.Engine) {
	api := server.Group("/api")

	api.GET("/temperature", getTemperature)
	api.GET("/humidity", getHumidity)

	server.GET("/health", healthCheck)
}
