package storage

import (
	"atom/container"
	"atom/providers/config"
	"log"

	"github.com/mojocn/base64Captcha"
)

func init() {
	if err := container.Container.Provide(NewCaptchaDriverDigit); err != nil {
		log.Fatal(err)
	}
}

func NewCaptchaDriverDigit(conf *config.Config) (base64Captcha.Driver, error) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	return base64Captcha.NewDriverDigit(
		int(conf.Captcha.ImgHeight),
		int(conf.Captcha.ImgWidth),
		int(conf.Captcha.KeyLong),
		0.7,
		80), nil
}
