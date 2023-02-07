package seeders

import (
	"atom/container"
	"atom/contracts"
	"atom/database/models"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(NewMigrationSeeder, dig.Group("seeders")); err != nil {
		log.Fatal(err)
	}
}

type MigrationSeeder struct {
}

func NewMigrationSeeder() contracts.Seeder {
	return &MigrationSeeder{}
}

func (s *MigrationSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	// times := 10
	// for i := 0; i < times; i++ {
	// 	data := s.Generate(faker, i)
	// 	if i == 0 {
	// 		stmt := &gorm.Statement{DB: db}
	// 		_ = stmt.Parse(&data)
	// 		log.Printf("seeding %s for %d times", stmt.Schema.Table, times)
	// 	}
	// 	db.Create(&data)
	// }
}

func (s *MigrationSeeder) Generate(faker *gofakeit.Faker, idx int) models.Migration {
	return models.Migration{
		// fill model fields
	}
}
