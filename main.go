package main

import (
	"github.com/kostya2011/dht-11-home-sensor/log"
)

func main() {
	// Initialize logging
	zapLogger := log.NewZapLogger()
	log.Init(zapLogger)

	// fmt.Println(config.Cfg.Log)
	// fmt.Println(config.Cfg.Server)

	log.Info("asdasd", log.SetField("b", "ddd"))
	log.Info("2211")
	// log.Error("sad")

	// routines.PrintLog()
}
