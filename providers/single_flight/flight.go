package single_flight

import (
	"atom/container"
	"log"

	"golang.org/x/sync/singleflight"
)

func init() {
	if err := container.Container.Provide(NewSingleFlight); err != nil {
		log.Fatal(err)
	}
}

func NewSingleFlight() (*singleflight.Group, error) {
	return &singleflight.Group{}, nil
}
