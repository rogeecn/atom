package mysql

import (
	"app/container"

	"gorm.io/gorm"
)

func init() {
	container.Container.Provide(NewMysqlConnection)
}

func NewMysqlConnection() (*gorm.DB, error) {
	return nil, nil
}
