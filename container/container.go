package container

import (
	"context"
	"log"

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
