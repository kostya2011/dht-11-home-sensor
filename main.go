package main

import (
	"fmt"

	"github.com/kostya2011/dht-11-home-sensor/config"
	"github.com/kostya2011/dht-11-home-sensor/log"
	"github.com/kostya2011/dht-11-home-sensor/routines"
)

func main() {
	// Initialize logging
	zapLogger := log.NewZapLogger()
	log.Init(zapLogger)

	fmt.Println(config.Cfg.Log)
	fmt.Println(config.Cfg.Server)

	log.Info("asdasd")
	log.Error("sad")

	routines.PrintLog()
}
