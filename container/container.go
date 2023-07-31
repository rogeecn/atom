package container

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/rogeecn/atom/utils/opt"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

var (
	Container *dig.Container = dig.New()
	Cancel    context.CancelFunc
	closeable []func()
)

func init() {
	closeable = make([]func(), 0)
	if err := Container.Provide(func() context.Context {
		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		go func() {
			<-ctx.Done()
			Close()
			cancel()
		}()
		Cancel = cancel
		return ctx
	}); err != nil {
		log.Fatal(err)
	}
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

func (p Providers) With(pcs ...Providers) Providers {
	for _, pc := range pcs {
		p = append(p, pc...)
	}
	return p
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
