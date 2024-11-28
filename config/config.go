package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/rogeecn/atom/container"
	"github.com/spf13/viper"
)

func Load(file, app string) (*viper.Viper, error) {
	v := viper.NewWithOptions(viper.KeyDelimiter("_"))
	v.AutomaticEnv()

	if file == "" {
		v.SetConfigType("toml")
		v.SetConfigName(app + ".toml")

		paths := []string{"."}
		// execute path
		execPath, err := os.Executable()
		if err == nil {
			paths = append(paths, filepath.Dir(execPath))
		}

		// home path
		homePath, err := os.UserHomeDir()
		if err == nil {
			paths = append(paths, homePath, homePath+"/"+app, homePath+"/.config", homePath+"/.config/"+app)
		}
		paths = append(paths, "/etc", "/etc/"+app, "/usr/local/etc", "/usr/local/etc/"+app)

		log.Println("try load config from paths:", paths)
		for _, path := range paths {
			v.AddConfigPath(path)
		}
	} else {
		v.SetConfigFile(file)
	}

	err := v.ReadInConfig()
	log.Println("use config file:", v.ConfigFileUsed())
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
