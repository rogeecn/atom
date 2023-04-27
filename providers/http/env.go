package http

import (
	"log"

	"github.com/spf13/viper"
)

const DefaultPrefix = "Http"

func AutoLoadConfig() *Config {
	return LoadConfig("")
}

func LoadConfig(file string) *Config {
	if file == "" {
		file = "config.toml"
	}

	v := viper.NewWithOptions(viper.KeyDelimiter("_"))
	v.AutomaticEnv()
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := v.UnmarshalKey(DefaultPrefix, &config); err != nil {
		log.Fatal(err)
	}

	return &config
}
