package database

import (
	"atom/providers/config"

	// "gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewSQLite(conf *config.SQLite) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(conf.File), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}
