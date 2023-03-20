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

type RoleService struct {
	dao  *dao.RoleDao
	uuid *uuid.Generator
}

func NewRoleService(
	dao *dao.RoleDao,
	uuid *uuid.Generator,
) *RoleService {
	return &RoleService{
		dao:  dao,
		uuid: uuid,
	}
}

func (svc *RoleService) GetByFilter(
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

func (svc *RoleService) treeMaker(ctx context.Context, models []*models.SysRole, pid uint64) []*dto.RoleTree {
	items := []*dto.RoleTree{}
	for _, model := range models {
		if model.ParentID == pid {
			items = append(items, &dto.RoleTree{
				ID:            model.ID,
				UUID:          model.UUID,
				Name:          model.Name,
				ParentID:      0,
				DefaultRouter: model.DefaultRouter,
				Children:      svc.treeMaker(ctx, models, model.ID),
			})
		}
	}
	return items
}

func (svc *RoleService) Tree(ctx context.Context) ([]*dto.RoleTree, error) {
	models, err := svc.dao.All(ctx)
	if err != nil {
		return nil, err
	}

	return svc.treeMaker(ctx, models, 0), nil
}

func (svc *RoleService) Create(ctx context.Context, req dto.RoleRequestForm) (*models.SysRole, error) {
	model := models.SysRole{
		UUID:          svc.uuid.MustGenerate(),
		Name:          req.Name,
		ParentID:      uint64(req.ParentID),
		DefaultRouter: req.DefaultRouter,
	}
	return svc.dao.Create(ctx, &model)
}
func (svc *RoleService) UpdateByID(ctx context.Context, id uint64, req dto.RoleRequestForm) (*models.SysRole, error) {
	model, err := svc.dao.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	model.Name = req.Name
	model.ParentID = uint64(req.ParentID)
	model.DefaultRouter = req.DefaultRouter

	return svc.dao.UpdateByID(ctx, model)
}

func (svc *RoleService) DeleteByID(ctx context.Context, id uint64) error {
	return svc.dao.DeleteByID(ctx, id)
}
