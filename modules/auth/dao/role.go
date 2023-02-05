package dao

import (
	"atom/common/request"
	"atom/database/models"
	"atom/database/query"
	"atom/modules/auth/dto"
	"context"
)

type RoleDao interface {
	GetByFilter(context.Context, dto.RoleRequestFilter, request.PageFilter) ([]*models.SysRole, uint64, error)
	FindByID(context.Context, uint64) (*models.SysRole, error)
	Create(context.Context, *models.SysRole) (*models.SysRole, error)
	UpdateByID(context.Context, *models.SysRole) (*models.SysRole, error)
	DeleteByID(context.Context, uint64) error
	DeletePermanentlyByID(context.Context, uint64) error
	All(context.Context) ([]*models.SysRole, error)
}

type roleDaoImpl struct {
	query *query.Query
}

func NewRoleDao(query *query.Query) RoleDao {
	return &roleDaoImpl{query: query}
}

func (dao *roleDaoImpl) GetByFilter(ctx context.Context, filter dto.RoleRequestFilter, page request.PageFilter) ([]*models.SysRole, uint64, error) {
	role := dao.query.SysRole
	query := role.WithContext(ctx)

	if filter.DefaultRouter != nil {
		query = query.Where(role.DefaultRouter.Eq(*filter.DefaultRouter))
	}

	if filter.Name != nil {
		query = query.Where(role.Name.Like(*filter.Name))
	}

	if filter.ParentID != nil {
		query = query.Where(role.ParentID.Eq(uint64(*filter.ParentID)))
	}

	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	items, err := query.Find()
	if err != nil {
		return nil, 0, err
	}

	return items, uint64(total), nil
}

func (dao *roleDaoImpl) All(ctx context.Context) ([]*models.SysRole, error) {
	role := dao.query.SysRole
	return role.WithContext(ctx).Find()
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
