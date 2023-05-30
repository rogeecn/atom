package hashids

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers/http"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "HashIDs"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(http.DefaultPrefix),
		},
	}
}

type Config struct {
	Salt      string
	MinLength uint
}
