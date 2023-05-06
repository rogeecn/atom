package atom

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers/grpcs"
	"github.com/rogeecn/atom/providers/log"
)

func DefaultGRPC(providers ...container.ProviderContainer) container.Providers {
	return append(container.Providers{
		log.DefaultProvider(),
		grpcs.DefaultProvider(),
	}, providers...)
}
