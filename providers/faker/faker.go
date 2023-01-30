package faker

import (
	"atom/container"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func init() {
	if err := container.Container.Provide(NewFaker); err != nil {
		log.Fatal(err)
	}
}

func NewFaker() *gofakeit.Faker {
	faker := gofakeit.New(time.Now().UnixNano())
	gofakeit.SetGlobalFaker(faker)

	return faker
}
