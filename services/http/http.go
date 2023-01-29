package http

import (
	_ "atom/modules"

	"atom/contracts"
	"atom/providers/config"
	"atom/providers/http"
	"atom/providers/logger"

	"go.uber.org/dig"
)

type Http struct {
	dig.In

	Conf    *config.Config
	Service *http.Service
	Routes  []contracts.Route `group:"route"`
}

func Serve(http Http) error {
	logger.Infof("http service port %s", http.Conf.Http.Address())
	for _, route := range http.Routes {
		route.Register()
	}

	logger.Infof("starting server on %s", http.Conf.Http.Address())
	return http.Service.Serve()
}
