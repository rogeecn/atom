package container

import (
	"context"

	"go.uber.org/dig"
)

var Container *dig.Container = dig.New()

func init() {
	Container.Provide(context.Background)
}
