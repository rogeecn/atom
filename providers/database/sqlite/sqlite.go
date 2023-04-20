package sqlite

import (
	"github.com/rogeecn/atom/container"
	"go.uber.org/dig"

	// "gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Provide(conf *Config, opts ...dig.ProvideOption) error {
	return container.Container.Provide(func() (*gorm.DB, error) {
		db, err := gorm.Open(sqlite.Open(conf.File), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		return db, err
	}, opts...)
}
