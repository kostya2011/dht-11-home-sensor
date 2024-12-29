package config

import "github.com/kostya2011/home-data/interfaces"

func OutputConfigString(ds interfaces.Outputter) {
	ds.Output()
}
