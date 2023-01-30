package contracts

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

// Migration route interface
type Seeder interface {
	Run(*gofakeit.Faker, *gorm.DB)
}
