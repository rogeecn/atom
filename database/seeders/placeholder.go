package seeders

import (
	"atom/container"
	"atom/contracts"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(NewPlaceholderSeeder, dig.Group("seeders")); err != nil {
		log.Fatal(err)
	}
}

type PlaceholderSeeder struct {
}

func NewPlaceholderSeeder() contracts.Seeder {
	return &PlaceholderSeeder{}
}

type Placeholder struct {
	gorm.Model

	Username string
}

func (s *PlaceholderSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {

}

func (s *PlaceholderSeeder) Generate(faker *gofakeit.Faker, idx int) Placeholder {
	return Placeholder{
		Username: faker.Name(),
	}
}
