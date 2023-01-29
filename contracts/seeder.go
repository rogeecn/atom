package contracts

import "gorm.io/gorm"

// Migration route interface
type Seeder interface {
	Times() uint
	Table() string
	Run(*gorm.DB)
}
