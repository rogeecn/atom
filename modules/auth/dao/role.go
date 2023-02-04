package dao

import (
	"atom/database/models"
	"atom/database/query"
	"context"
)

type RoleDao interface {
	FindByID(context.Context, uint64) (*models.SysRole, error)
	Create(context.Context, *models.SysRole) (*models.SysRole, error)
	UpdateByID(context.Context, *models.SysRole) (*models.SysRole, error)
	DeleteByID(context.Context, uint64) error
	DeletePermanentlyByID(context.Context, uint64) error
}

type roleDaoImpl struct {
	query *query.Query
}

func NewRoleDao(query *query.Query) RoleDao {
	return &roleDaoImpl{query: query}
}

func (dao *roleDaoImpl) FindByID(ctx context.Context, id uint64) (*models.SysRole, error) {
	role := dao.query.SysRole
	return role.WithContext(ctx).Where(role.ID.Eq(id)).First()
}

func (dao *roleDaoImpl) Create(ctx context.Context, model *models.SysRole) (*models.SysRole, error) {
	role := dao.query.SysRole
	if err := role.WithContext(ctx).Create(model); err != nil {
		return nil, err
	}
	return model, nil
}

func (dao *roleDaoImpl) UpdateByID(ctx context.Context, model *models.SysRole) (*models.SysRole, error) {
	role := dao.query.SysRole
	_, err := role.WithContext(ctx).Where(role.ID.Eq(model.ID)).Updates(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (dao *roleDaoImpl) DeleteByID(ctx context.Context, id uint64) error {
	role := dao.query.SysRole
	_, err := role.WithContext(ctx).Where(role.ID.Eq(id)).Delete()
	return err
}

func (dao *roleDaoImpl) DeletePermanentlyByID(ctx context.Context, id uint64) error {
	role := dao.query.SysRole
	_, err := role.WithContext(ctx).Unscoped().Where(role.ID.Eq(id)).Delete()
	return err
}
