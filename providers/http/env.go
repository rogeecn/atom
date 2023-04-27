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

	v := viper.NewWithOptions(viper.KeyDelimiter("_"))

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	if !fs.FileExist(file) {
		return &Config{}
	}

	// load file
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return &config
}
