package single_flight

import (
	"github.com/rogeecn/atom/container"
	"go.uber.org/dig"
	"golang.org/x/sync/singleflight"
)

func Provide(opts ...dig.ProvideOption) error {
	return container.Container.Provide(func() (*singleflight.Group, error) {
		return &singleflight.Group{}, nil
	}, opts...)
}
