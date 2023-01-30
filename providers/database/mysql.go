package mysql

import (
	"atom/container"
	"atom/providers/config"
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(NewDatabase); err != nil {
		log.Fatal(err)
	}
}

func NewDatabase(config *config.Config) (*gorm.DB, error) {
	sqlDB, err := sql.Open("mysql", config.Database.MySQL.DSN())
	if err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return gormDB, err
}
