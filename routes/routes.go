package routes

import (
	"github.com/gin-gonic/gin"
)

func NewRoutes(server *gin.Engine) {
	api := server.Group("/api")

	api.GET("/ticker", getTickerValue)
	api.GET("/tickers", getTickerValue)
	api.POST("/ticker/register", registerTicker)
	api.DELETE("/ticker/unregister", registerTicker)

	server.GET("/health", healthCheck)
}
