package database

import (
	"atom/container"
	"atom/providers/config"
	"errors"
	"log"

	"gorm.io/gorm"
)

const (
	DriverMySQL     = "mysql"
	DriverSQLite    = "sqlite"
	DriverPostgres  = "postgres"
	DriverSQLServer = "sqlserver"
)

func init() {
	if err := container.Container.Provide(NewDatabase); err != nil {
		log.Fatal(err)
	}
}

func NewDatabase(config *config.Config) (*gorm.DB, error) {
	switch config.Database.Driver {
	case DriverMySQL:
		return NewMySQL(config.Database.MySQL)
	case DriverSQLite:
		return NewSQLite(config.Database.SQLite)
	case DriverPostgres:
		return NewPostgres(config.Database.PostgreSQL)
	}
	return nil, errors.New("failed to connect to db")
}
