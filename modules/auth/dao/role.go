package dao

import (
	"atom/common/request"
	"atom/database/models"
	"atom/database/query"
	"atom/modules/auth/dto"
	"context"
)

type RoleDao struct {
	query *query.Query
}

func NewRoleDao(query *query.Query) *RoleDao {
	return &RoleDao{query: query}
}

func (dao *RoleDao) GetByFilter(ctx context.Context, filter dto.RoleRequestFilter, page request.PageFilter) ([]*models.SysRole, uint64, error) {
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

func (dao *RoleDao) All(ctx context.Context) ([]*models.SysRole, error) {
	role := dao.query.SysRole
	return role.WithContext(ctx).Find()
}

func (dao *RoleDao) FindByID(ctx context.Context, id uint64) (*models.SysRole, error) {
	role := dao.query.SysRole
	return role.WithContext(ctx).Where(role.ID.Eq(id)).First()
}

func (dao *RoleDao) Create(ctx context.Context, model *models.SysRole) (*models.SysRole, error) {
	role := dao.query.SysRole
	if err := role.WithContext(ctx).Create(model); err != nil {
		return nil, err
	}
	return model, nil
}

func (dao *RoleDao) UpdateByID(ctx context.Context, model *models.SysRole) (*models.SysRole, error) {
	role := dao.query.SysRole
	_, err := role.WithContext(ctx).Where(role.ID.Eq(model.ID)).Updates(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (dao *RoleDao) DeleteByID(ctx context.Context, id uint64) error {
	role := dao.query.SysRole
	_, err := role.WithContext(ctx).Where(role.ID.Eq(id)).Delete()
	return err
}

func (dao *RoleDao) DeletePermanentlyByID(ctx context.Context, id uint64) error {
	role := dao.query.SysRole
	_, err := role.WithContext(ctx).Unscoped().Where(role.ID.Eq(id)).Delete()
	return err
}
