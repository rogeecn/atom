package jwt

import (
	"time"

	"github.com/rogeecn/atom/providers/log"
)

const DefaultPrefix = "JWT"

type Config struct {
	SigningKey  string // jwt签名
	ExpiresTime string // 过期时间
	Issuer      string // 签发者
}

func (c *Config) ExpiresTimeDuration() time.Duration {
	d, err := time.ParseDuration(c.ExpiresTime)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
