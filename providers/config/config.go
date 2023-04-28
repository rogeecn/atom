package config

import (
	"github.com/rogeecn/atom/container"
	"github.com/spf13/viper"
)

func Load(file string) (*viper.Viper, error) {
	v := viper.NewWithOptions(viper.KeyDelimiter("_"))
	v.AutomaticEnv()
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	err := container.Container.Provide(func() (*viper.Viper, error) {
		return v, nil
	})
	if err != nil {
		return nil, err
	}

	return v, nil
}
