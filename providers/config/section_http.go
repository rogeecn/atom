package config

import "fmt"

type Http struct {
	Static    string
	Host      string
	Port      uint
	Https     bool
	HttpsCert string
	HttpKey   string
}

func (h Http) Address() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

func (h Http) PortString() string {
	return fmt.Sprintf(":%d", h.Port)
}
