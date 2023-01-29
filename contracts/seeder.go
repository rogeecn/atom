package contracts

import "gorm.io/gorm"

// Migration route interface
type Seeder interface {
	Run(*gorm.DB)
}
