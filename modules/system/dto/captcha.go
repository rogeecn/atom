package dto

type SysCaptchaResponse struct {
	CaptchaId     string `json:"captcha_id,omitempty"`
	PicPath       string `json:"pic_path,omitempty"`
	CaptchaLength int    `json:"captcha_length,omitempty"`
	OpenCaptcha   bool   `json:"open_captcha,omitempty"`
}
