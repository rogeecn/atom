package jwt

import (
	"time"

	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers/http"
	"github.com/rogeecn/atom/providers/log"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "JWT"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(http.DefaultPrefix),
		},
	}
}

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
