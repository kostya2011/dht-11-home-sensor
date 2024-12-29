package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kostya2011/home-data/config"
	"github.com/kostya2011/home-data/log"
)

func newGin() *gin.Engine {
	cfg := config.Values.Server

	switch {
	case cfg.Mode == "release":
		gin.SetMode(gin.ReleaseMode)
	case cfg.Mode == "test":
		gin.SetMode(gin.TestMode)
	case cfg.Mode == "debug":
		gin.SetMode(gin.DebugMode)
	default:
		log.Warn("Unknown mode selected. Falling back to `Release mode`")
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Debug("endpoint", log.SetField("method", httpMethod))
	}

	return gin.Default()
}
