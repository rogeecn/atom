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
	if err := container.Container.Provide(NewSysRoleSeeder, dig.Group("seeders")); err != nil {
		log.Fatal(err)
	}
}

type SysRoleSeeder struct {
}

func NewSysRoleSeeder() contracts.Seeder {
	return &SysRoleSeeder{}
}

func (s *SysRoleSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	times := 50
	for i := 0; i < times; i++ {
		data := s.Generate(faker, i)
		if i == 0 {
			stmt := &gorm.Statement{DB: db}
			_ = stmt.Parse(&data)
			log.Printf("seeding %s for %d times", stmt.Schema.Table, times)
		}
		db.Create(&data)
	}
}

func (s *SysRoleSeeder) Generate(faker *gofakeit.Faker, idx int) models.SysRole {
	parentID := 0
	if idx > 10 {
		parentID = faker.IntRange(1, 10)
	}

	if idx > 20 {
		parentID = faker.IntRange(10, 20)
	}

	if idx > 30 {
		parentID = faker.IntRange(20, 30)
	}

	if idx > 40 {
		parentID = faker.IntRange(30, 40)
	}

	return models.SysRole{
		UUID:          faker.UUID(),
		Name:          faker.Name(),
		ParentID:      uint64(parentID),
		DefaultRouter: faker.Name(),
	}
}
