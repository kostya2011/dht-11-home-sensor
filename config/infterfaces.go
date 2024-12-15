package config

type outputter interface {
	Display()
}

func DisplayConfig(config outputter) {
	config.Display()
}
