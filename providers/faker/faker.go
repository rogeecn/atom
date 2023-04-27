package faker

import (
	"time"

	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers"

	"github.com/brianvoe/gofakeit/v6"
)

func Provide(o *providers.Options) error {
	return container.Container.Provide(func() (*gofakeit.Faker, error) {
		faker := gofakeit.New(time.Now().UnixNano())
		gofakeit.SetGlobalFaker(faker)

		return faker, nil
	}, o.DiOptions()...)
}
