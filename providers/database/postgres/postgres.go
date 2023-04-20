package postgres

import (
	"log"

	"github.com/rogeecn/atom/container"
	"go.uber.org/dig"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Provide(conf *Config, opts ...dig.ProvideOption) error {
	return container.Container.Provide(func() (*gorm.DB, error) {
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
	}, opts...)
}
