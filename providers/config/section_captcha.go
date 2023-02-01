package config

import (
	"log"
	"time"
)

type Captcha struct {
	KeyLong            int    // 验证码长度
	ImgWidth           int    // 验证码宽度
	ImgHeight          int    // 验证码高度
	OpenCaptcha        int    // 防爆破验证码开启此数，0代表每次登录都需要验证码，其他数字代表错误密码此数，如3代表错误三次后出现验证码
	OpenCaptchaTimeOut string // 防爆破验证码超时时间，单位：s(秒)
}

func (c *Captcha) OpenCaptchaTimeOutDuration() time.Duration {
	d, err := time.ParseDuration(c.OpenCaptchaTimeOut)
	if err != nil {
		log.Panic(err)
	}
	return d
}
