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
var closeable []func()

func init() {
	if err := Container.Provide(func() context.Context {
		ctx, cancel := context.WithCancel(context.Background())
		Cancel = cancel
		return ctx
	}); err != nil {
		log.Fatal(err)
	}

	closeable = make([]func(), 0)
}

func AddCloseAble(c func()) {
	closeable = append(closeable, c)
}

func Close() {
	for _, c := range closeable {
		c()
	}
}

type ProviderContainer struct {
	Provider func(...opt.Option) error
	Options  []opt.Option
}

type Providers []ProviderContainer

func (p Providers) With(pcs Providers) Providers {
	return append(p, pcs...)
}

func (p Providers) Provide(config *viper.Viper) error {
	for _, provider := range p {
		provider.Options = append(provider.Options, opt.Config(config))
		if err := provider.Provider(provider.Options...); err != nil {
			return err
		}
	}
	return nil
}
