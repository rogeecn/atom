package http

import (
	"log"

	"github.com/rogeecn/atom/utils/fs"
	"github.com/spf13/viper"
)

const DefaultPrefix = "HTTP"

func AutoLoadConfig() *Config {
	return LoadConfig("", DefaultPrefix)
}

func LoadConfig(file, envPrefix string) *Config {
	if file == "" {
		file = "config.toml"
	}

	if envPrefix == "" {
		envPrefix = DefaultPrefix
	}

	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	if !fs.FileExist(file) {
		return &Config{}
	}

	// load file
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return &config
}
