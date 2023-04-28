package container

import (
	"context"
	"log"

	"github.com/rogeecn/atom/utils/opt"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

var Container *dig.Container = dig.New()
var Cancel context.CancelFunc

func init() {
	if err := Container.Provide(func() context.Context {
		ctx, cancel := context.WithCancel(context.Background())
		Cancel = cancel
		return ctx
	}); err != nil {
		log.Fatal(err)
	}
}

type ProviderContainer struct {
	Provider func(...opt.Option) error
	Options  []opt.Option
}

type Providers []ProviderContainer

func (p Providers) Provide(config *viper.Viper) error {
	for _, provider := range p {
		provider.Options = append(provider.Options, opt.Config(config))
		if err := provider.Provider(provider.Options...); err != nil {
			return err
		}
	}
	return nil
}
