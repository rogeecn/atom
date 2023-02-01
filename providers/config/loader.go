package config

import (
	"atom/container"
	"atom/utils"
	"log"

	"github.com/pkg/errors"
	"github.com/rogeecn/fabfile"
	"github.com/spf13/viper"
)

type Config struct {
	App      App
	Captcha  Captcha
	Http     Http
	Log      Log
	Database Database
	Storage  Storage
}

func init() {
	if err := container.Container.Provide(Load); err != nil {
		log.Fatal(err)
	}
}

func Load() (*Config, error) {
	var err error
	confFile := utils.ShareConfigFile
	if confFile == "" {
		confFile, err = fabfile.Find("config.toml")
		if err != nil {
			return nil, err
		}
	}

	viper.SetConfigName(confFile) // name of config file (without extension)
	viper.SetConfigType("toml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/") // call multiple times to add many search paths
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		return nil, errors.Wrapf(err, "read config failed, %s", confFile)
	}

	config := &Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.Wrapf(err, "unmarshal data failed, %s", confFile)
	}

	return config, nil

}
