package database

import (
	"atom/providers/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewPostgres(conf *config.PostgreSQL) (*gorm.DB, error) {
	dbConfig := postgres.Config{
		DSN: conf.DSN(), // DSN data source name
	}
	log.Println("PostgreSQL DSN: ", dbConfig.DSN)

	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.Prefix,
			SingularTable: conf.Singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	db, err := gorm.Open(postgres.New(dbConfig), &gormConfig)
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)

	return db, err
}
