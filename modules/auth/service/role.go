package service

import (
	"atom/common/request"
	"atom/common/response"
	"atom/database/models"
	"atom/modules/auth/dao"
	"atom/modules/auth/dto"
	"atom/providers/uuid"
	"context"
)

type RoleService interface {
	GetByFilter(context.Context, dto.RoleRequestFilter, request.PageFilter) (*response.PageResponse[*models.SysRole], error)
	Create(context.Context, dto.RoleRequestForm) (*models.SysRole, error)
	UpdateByID(context.Context, uint64, dto.RoleRequestForm) (*models.SysRole, error)
	DeleteByID(context.Context, uint64) error
}

type roleService struct {
	dao  dao.RoleDao
	uuid *uuid.Generator
}

func NewRoleService(
	dao dao.RoleDao,
	uuid *uuid.Generator,
) RoleService {
	return &roleService{
		dao:  dao,
		uuid: uuid,
	}
}

func (svc *roleService) GetByFilter(
	ctx context.Context,
	filter dto.RoleRequestFilter,
	page request.PageFilter,
) (*response.PageResponse[*models.SysRole], error) {
	items, count, err := svc.dao.GetByFilter(ctx, filter, page)
	if err != nil {
		return nil, err
	}

	return &response.PageResponse[*models.SysRole]{
		Items: items,
		Total: count,
	}, nil
}

func (svc *roleService) Create(ctx context.Context, req dto.RoleRequestForm) (*models.SysRole, error) {
	model := models.SysRole{
		UUID:          svc.uuid.MustGenerate(),
		Name:          req.Name,
		ParentID:      uint64(req.ParentID),
		DefaultRouter: req.DefaultRouter,
	}
	return svc.dao.Create(ctx, &model)
}
func (svc *roleService) UpdateByID(ctx context.Context, id uint64, req dto.RoleRequestForm) (*models.SysRole, error) {
	model, err := svc.dao.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	model.Name = req.Name
	model.ParentID = uint64(req.ParentID)
	model.DefaultRouter = req.DefaultRouter

	return svc.dao.UpdateByID(ctx, model)
}

func (svc *roleService) DeleteByID(ctx context.Context, id uint64) error {
	return svc.dao.DeleteByID(ctx, id)
}
