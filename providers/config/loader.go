package config

import (
	"atom/container"
	"atom/utils"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-micro/plugins/v4/config/encoder/toml"
	"github.com/go-micro/plugins/v4/config/source/etcd"
	"github.com/rogeecn/fabfile"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
	"go-micro.dev/v4/config/source/env"
	"go-micro.dev/v4/config/source/file"
	"go-micro.dev/v4/logger"
)

var c *Config

type Config struct {
	App  App
	Http Http
	Log  Log
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

	options := []config.Option{}

	options = append(options, config.WithSource(file.NewSource(
		file.WithPath(confFile),
		source.WithEncoder(toml.NewEncoder()),
	)))

	etcdEndpoints := etcdEndpoints()
	if len(etcdEndpoints) > 0 {
		logger.Info("etcd endpoints: ", etcdEndpoints, len(etcdEndpoints))
		options = append(options, config.WithSource(etcd.NewSource(
			etcd.WithAddress(etcdEndpoints...),
			etcd.WithPrefix("/micro/config/api.web"),
			etcd.StripPrefix(true),
		)))
	}

	options = append(options, config.WithSource(env.NewSource()))

	options = append(options, config.WithReader(json.NewReader(reader.WithEncoder(toml.NewEncoder()))))

	conf, err := config.NewConfig(options...)
	if err != nil {
		return nil, err
	}

	if err := conf.Scan(&c); err != nil {
		return nil, err
	}

	go func() {
		ticker := time.NewTicker(time.Second * 10)
		defer ticker.Stop()

		for range ticker.C {
			if err := conf.Scan(&c); err != nil {
				logger.Fatal(err)
			}
		}
	}()
	return c, nil
}

func etcdEndpoints() []string {
	var endpoints []string
	envVars := strings.Split(os.Getenv("ETCD_ENDPOINTS"), ",")
	for _, env := range envVars {
		if strings.TrimSpace(env) == "" {
			continue
		}
		endpoints = append(endpoints, strings.TrimSpace(env))
	}
	return endpoints
}
