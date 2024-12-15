package main

import (
	"github.com/kostya2011/dht-11-home-sensor/log"
)

func main() {
	// Initialize logging
	// logger := zap.Must(zap.NewProduction())
	// defer logger.Sync()
	log.Debug("sad", log.IntLogField("dsa", 4))
	log.Info("sad2", log.IntLogField("dsa", 4))
	log.Warn("sad3", log.IntLogField("dsa", 4))
	log.Erorr("sad4", log.IntLogField("dsa", 4))
	log.Fatal("sad5", log.StringLogField("dsa2", "dd"))
	log.Panic("sad6", log.StringLogField("dsa2", "dd"))
	// fmt.Println(config.Cfg.Log)
	// fmt.Println(config.Cfg.Server)
}
