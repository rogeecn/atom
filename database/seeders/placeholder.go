package seeders

import (
	"atom/container"
	"atom/contracts"
	"log"

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
}

func (s *PlaceholderSeeder) Run(db *gorm.DB) {

}

func (s *PlaceholderSeeder) Generate(idx int) Placeholder {
	return Placeholder{}
}
