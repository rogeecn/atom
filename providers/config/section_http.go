package config

import (
	"fmt"
	"log"
	"time"
)

type Http struct {
	Static    string
	Host      string
	Port      uint
	Https     bool
	HttpsCert string
	HttpKey   string
	Cors      struct {
		Mode      string
		Whitelist []Whitelist
	}
	JWT JWT
}

type JWT struct {
	SigningKey  string // jwt签名
	ExpiresTime string // 过期时间
	BufferTime  string // 缓冲时间
	Issuer      string // 签发者
}

func (j JWT) ExpiresTimeDuration() time.Duration {
	d, err := time.ParseDuration(j.ExpiresTime)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

func (j JWT) BufferTimeDuration() time.Duration {
	d, err := time.ParseDuration(j.BufferTime)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

type Whitelist struct {
	AllowOrigin      string
	AllowHeaders     string
	AllowMethods     string
	ExposeHeaders    string
	AllowCredentials bool
}

func (h Http) Address() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

func (h Http) PortString() string {
	return fmt.Sprintf(":%d", h.Port)
}
