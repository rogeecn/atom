package contracts

import "gorm.io/gorm"

// Migration route interface
type Migration interface {
	ID() string
	Up(tx *gorm.DB) error
	Down(tx *gorm.DB) error
}
