package captcha

import (
	"atom/container"
	"atom/providers/config"
	"errors"
	"log"

	"github.com/mojocn/base64Captcha"
)

func init() {
	if err := container.Container.Provide(NewCaptcha); err != nil {
		log.Fatal(err)
	}
}

type CaptchaResponse struct {
	CaptchaId     string `json:"captcha_id,omitempty"`
	PicPath       string `json:"pic_path,omitempty"`
	CaptchaLength uint   `json:"captcha_length,omitempty"`
	OpenCaptcha   uint   `json:"open_captcha,omitempty"`
}

type Captcha struct {
	conf    *config.Config
	captcha *base64Captcha.Captcha
}

func NewCaptcha(conf *config.Config, driver base64Captcha.Driver) (*Captcha, error) {
	var store = base64Captcha.DefaultMemStore
	return &Captcha{
		conf:    conf,
		captcha: base64Captcha.NewCaptcha(driver, store),
	}, nil

}

func (c *Captcha) Generate() (*CaptchaResponse, error) {
	id, b64s, err := c.captcha.Generate()
	if err != nil {
		return nil, errors.New("验证码获取失败")
	}

	return &CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: c.conf.Captcha.KeyLong,
		OpenCaptcha:   c.conf.Captcha.OpenCaptcha,
	}, nil
}

func (c *Captcha) Verify(id, answer string) bool {
	return c.captcha.Verify(id, answer, false)
}
