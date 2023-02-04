package controller

import (
	"atom/providers/captcha"
	"atom/providers/config"

	"github.com/gin-gonic/gin"
)

type CaptchaController interface {
	Show(*gin.Context) (*captcha.CaptchaResponse, error)
}

type captchaControllerImpl struct {
	conf    *config.Config
	captcha *captcha.Captcha
}

func NewCaptchaController(
	conf *config.Config,
	captcha *captcha.Captcha,
) CaptchaController {
	return &captchaControllerImpl{
		conf:    conf,
		captcha: captcha,
	}
}

func (c *captchaControllerImpl) Show(ctx *gin.Context) (*captcha.CaptchaResponse, error) {
	return c.captcha.Generate()
}
