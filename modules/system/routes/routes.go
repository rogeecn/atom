package routes

import (
	"atom/contracts"
	"atom/modules/system/controller"
	"atom/providers/http"

	"github.com/rogeecn/gen"
)

type Route struct {
	captcha *controller.CaptchaController
	svc     *http.Service
}

func NewRoute(captcha *controller.CaptchaController, svc *http.Service) contracts.Route {
	return &Route{captcha: captcha, svc: svc}
}

func (r *Route) Register() {
	r.svc.Engine.GET("/captcha", gen.DataFunc(r.captcha.Show))
}
