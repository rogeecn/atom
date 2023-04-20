package http

import (
	"fmt"
)

type Config struct {
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
}

type Whitelist struct {
	AllowOrigin      string
	AllowHeaders     string
	AllowMethods     string
	ExposeHeaders    string
	AllowCredentials bool
}

func (h *Config) Address() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

func (h *Config) PortString() string {
	return fmt.Sprintf(":%d", h.Port)
}
