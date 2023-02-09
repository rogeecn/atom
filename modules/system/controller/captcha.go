package controller

import (
	"atom/providers/captcha"
	"atom/providers/config"

	"github.com/gin-gonic/gin"
)

type CaptchaController struct {
	conf    *config.Config
	captcha *captcha.Captcha
}

func NewCaptchaController(
	conf *config.Config,
	captcha *captcha.Captcha,
) *CaptchaController {
	return &CaptchaController{
		conf:    conf,
		captcha: captcha,
	}
}

func (c *CaptchaController) Show(ctx *gin.Context) (*captcha.CaptchaResponse, error) {
	return c.captcha.Generate()
}
