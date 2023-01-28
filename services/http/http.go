package http

import (
	_ "app/modules"

	"app/contracts"
	"app/providers/config"
	"app/providers/http"
	"app/providers/logger"

	"go.uber.org/dig"
)

type Http struct {
	dig.In

	Logger  *logger.Logger
	Conf    *config.Config
	Service *http.Service
	Routes  []contracts.Route `group:"route"`
}

func Serve(http Http) error {
	http.Logger.Infof("http service port %s", http.Conf.Http.Address())
	for _, route := range http.Routes {
		route.Register()
	}

	http.Logger.Infof("starting server on %s", http.Conf.Http.Address())
	return http.Service.Serve()
}
