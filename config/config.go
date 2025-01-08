package config

import (
	"log"
	"path/filepath"

	"git.ipao.vip/rogeecn/atom/container"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func Load(file string) (*viper.Viper, error) {
	v := viper.NewWithOptions(viper.KeyDelimiter("_"))
	v.AutomaticEnv()

	ext := filepath.Ext(file)
	if ext == "" {
		v.SetConfigType("toml")
		v.SetConfigFile(file)
	} else {
		v.SetConfigType(ext[1:])
		v.SetConfigFile(file)
	}

	v.AddConfigPath(".")

	err := v.ReadInConfig()
	log.Println("config file:", v.ConfigFileUsed())
	if err != nil {
		return nil, errors.Wrap(err, "config file read error")
	}

	err = container.Container.Provide(func() (*viper.Viper, error) {
		return v, nil
	})
	if err != nil {
		return nil, err
	}

	return v, nil
}
