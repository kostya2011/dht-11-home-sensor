package main

import (
	"github.com/kostya2011/home-data/config"
	"github.com/kostya2011/home-data/log"
	"github.com/kostya2011/home-data/routines"
)

func main() {
	// Initialize logging
	zapLogger := log.NewZapLogger()
	log.Init(zapLogger)

	log.Debug(config.Values.Output())

	// fmt.Println(config.Cfg.Log)
	// fmt.Println(config.Cfg.Server)

	log.Info("asdasd", log.SetField("b", "ddd"))
	log.Info("2211")
	// log.Error("sad")

	routines.PrintLog()
}
