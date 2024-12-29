package main

import (
	"fmt"

	"github.com/kostya2011/home-data/config"
	"github.com/kostya2011/home-data/log"
	"github.com/kostya2011/home-data/routes"
)

func main() {
	// Initialize logging
	cfg := config.Values
	switch {
	case cfg.Log.Logger == "zap":
		zapLogger := log.NewZapLogger()
		log.Init(zapLogger)
	default:
		fmt.Println("Falling back to default looger `zap`")
		zapLogger := log.NewZapLogger()
		log.Init(zapLogger)
	}

	log.Info("Initilize Gin")
	server := newGin()
	routes.NewRoutes(server)

	server.Run(fmt.Sprintf("0.0.0.0:%v", cfg.Server.Port))
}
