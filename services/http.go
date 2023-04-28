package services

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers/http"
	"go.uber.org/dig"
)

type Http struct {
	dig.In

	Service http.Service
	Routes  []http.Route `group:"routes"`
}

func ServeHttp() error {
	return container.Container.Invoke(func(http Http) error {
		for _, route := range http.Routes {
			route.Register()
		}

		return http.Service.Serve()
	})
}
