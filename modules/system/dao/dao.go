package dao

import (
	"atom/providers/config"
	"context"
	"errors"

	"gorm.io/gorm"
)

type Dao struct {
	Conf *config.Config
	DB   *gorm.DB
}

func NewDao(DB *gorm.DB) *Dao {
	return &Dao{DB: DB}
}

func (c *Dao) Release(ctx context.Context, a int, b string) error {
	if a == 20 {
		return errors.New("A cant't be 20 ")
	}
	return nil
}
