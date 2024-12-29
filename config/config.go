package config

import (
	"errors"
	"sync"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	once   sync.Once
	Values *Config
)

type ServerConfig struct {
	Mode string `mapstructure:"mode"`
	Port string `mapstructure:"port"`
}

type LogConfig struct {
	Logger          string   `mapstructure:"logger"`
	Level           string   `mapstructure:"level"`
	EncondingFormat string   `mapstructure:"encodingFormat"`
	Outputs         []string `mapstructure:"outputs"`
	ErrorOutputs    []string `mapstructure:"errorOutputs"`
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
	viper.SetDefault("server.mode", "dev")
	viper.SetDefault("server.port", "8080")

	// Set defaults for logs config
	viper.SetDefault("log.level", "error")
	viper.SetDefault("log.logger", "zap")
	viper.SetDefault("Log.encodingFormat", "json")
	viper.SetDefault("log.outputs", []string{"stdout"})
	viper.SetDefault("log.errorOutputs", []string{"stderr"})
}

func (cfg *Config) Output() string {
	strConfig, _ := yaml.Marshal(cfg)
	return string(strConfig)
}

func init() {
	once.Do(func() {
		// Read config
		if err := readConfig(); err != nil {
			panic(err)
		}

		// Unmarshall config into Config structure
		Values = &Config{}
		setDefaults()

		if err := viper.Unmarshal(Values); err != nil {
			panic(err)
		}
	})

}
