package single_flight

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers"
	"golang.org/x/sync/singleflight"
)

func Provide(o *providers.Options) error {
	return container.Container.Provide(func() (*singleflight.Group, error) {
		return &singleflight.Group{}, nil
	}, o.DiOptions()...)
}
