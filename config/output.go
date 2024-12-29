package config

import "github.com/kostya2011/home-data/interfaces"

func Output(ds interfaces.Outputter) {
	ds.Output()
}
