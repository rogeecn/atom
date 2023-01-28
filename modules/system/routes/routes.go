package routes

import (
	"atom/contracts"
	"atom/modules/system/controller"
	"atom/providers/http"

	"github.com/rogeecn/gen"
)

type Route struct {
	controller controller.Controller
	svc        *http.Service
}

func NewRoute(c controller.Controller, svc *http.Service) contracts.Route {
	return &Route{controller: c, svc: svc}
}

func (r *Route) Register() {
	r.svc.Engine.GET("/name", gen.DataFunc(r.controller.GetName))
}
