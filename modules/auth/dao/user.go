package dao

import (
	"atom/database/models"
	"atom/database/query"
	"context"
)

type UserDao struct {
	query *query.Query
}

func NewUserDao(query *query.Query) *UserDao {
	return &UserDao{query: query}
}

func (dao *UserDao) FindByID(ctx context.Context, id uint64) (*models.User, error) {
	user := dao.query.User
	return user.WithContext(ctx).Where(user.ID.Eq(id)).First()
}

func (dao *UserDao) Create(ctx context.Context, model *models.User) (*models.User, error) {
	user := dao.query.User
	if err := user.WithContext(ctx).Create(model); err != nil {
		return nil, err
	}
	return model, nil
}
