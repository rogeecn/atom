package faker

import (
	"time"

	"github.com/rogeecn/atom/container"
	"go.uber.org/dig"

	"github.com/brianvoe/gofakeit/v6"
)

func Provide(opts ...dig.ProvideOption) error {
	return container.Container.Provide(func() (*gofakeit.Faker, error) {
		faker := gofakeit.New(time.Now().UnixNano())
		gofakeit.SetGlobalFaker(faker)

		return faker, nil
	}, opts...)
}
