package controller

import (
	"atom/modules/system/dto"
	"atom/providers/config"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type CaptchaController interface {
	Show(*gin.Context) (dto.SysCaptchaResponse, error)
}

type captchaControllerImpl struct {
	conf *config.Config
}

func NewCaptchaController(conf *config.Config) CaptchaController {
	return &captchaControllerImpl{conf: conf}
}

func (c *captchaControllerImpl) Show(ctx *gin.Context) (dto.SysCaptchaResponse, error) {
	// 判断验证码是否开启
	var store = base64Captcha.DefaultMemStore

	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(c.conf.Captcha.ImgHeight, c.conf.Captcha.ImgWidth, c.conf.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		return dto.SysCaptchaResponse{}, errors.New("验证码获取失败")
	}

	return dto.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: c.conf.Captcha.KeyLong,
		OpenCaptcha:   c.conf.Captcha.OpenCaptcha != 0,
	}, nil
}
