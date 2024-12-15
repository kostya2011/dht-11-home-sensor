package config

import (
	"errors"
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	Cfg  *Config
)

type ServerConfig struct {
	Mode string `mapstructure:"mode"`
	Port string `mapstructure:"port"`
}

type LogConfig struct {
	Level           string   `mapstructure:"level"`
	Outputs         []string `mapstructure:"outputs"`
	ErrorOutputs    []string `mapstructure:"errorOutputs"`
	EncondingFormat string   `mapstructure:"encodingFormat"`
}

type Config struct {
	Server *ServerConfig `mapstructure:"server"`
	Log    *LogConfig    `mapstructure:"log"`
}

func readConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("properties") // Register config file name (no extension)
	viper.SetConfigType("yml")        // Look for specific type
	viper.ReadInConfig()

	if err := viper.ReadInConfig(); err != nil {
		return errors.New("Could not read config file.")
	}
	return nil
}

func setDefaults() {
	// Set defaults for server config
	viper.SetDefault("Server.Mode", "dev")
	viper.SetDefault("Server.Port", "8080")

	// Set defaults for logs config
	viper.SetDefault("Log.Level", "error")
	viper.SetDefault("Log.EncondingFormat", "json")
	viper.SetDefault("Log.Outputs", []string{"stdout"})
	viper.SetDefault("Log.ErrorOutputs", []string{"stderr"})
}

func init() {
	once.Do(func() {
		// Read config
		if err := readConfig(); err != nil {
			panic(err)
		}

		// Unmarshall config into Config structure
		Cfg = &Config{}
		setDefaults()

		if err := viper.Unmarshal(Cfg); err != nil {
			panic(err)
		}
	})

}
