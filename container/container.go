package container

import (
	"context"
	"log"

	"go.uber.org/dig"
)

var Container *dig.Container = dig.New()

func init() {
	if err := Container.Provide(context.Background); err != nil {
		log.Fatal(err)
	}
}
