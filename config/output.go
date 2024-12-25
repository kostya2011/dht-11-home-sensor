package config

import "github.com/kostya2011/dht-11-home-sensor/interfaces"

func Output(ds interfaces.Outputter) {
	ds.Output()
}
