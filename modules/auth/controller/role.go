package controller

import (
	"atom/common/request"
	"atom/common/response"
	"atom/database/models"
	"atom/modules/auth/dto"
	"atom/modules/auth/service"

	"github.com/gin-gonic/gin"
)

type RoleController interface {
	GetByFilter(*gin.Context, dto.RoleRequestFilter, request.PageFilter) (*response.PageResponse[*models.SysRole], error)
	Tree(*gin.Context) ([]*dto.RoleTree, error)
	Create(*gin.Context, dto.RoleRequestForm) error
	Delete(*gin.Context, int) error
	UpdateByID(*gin.Context, int, dto.RoleRequestForm) error
}

type roleControllerImpl struct {
	roleSvc service.RoleService
}

func NewRoleController(
	roleSvc service.RoleService,
) RoleController {
	return &roleControllerImpl{
		roleSvc: roleSvc,
	}
}

func (c *roleControllerImpl) GetByFilter(
	ctx *gin.Context,
	filter dto.RoleRequestFilter,
	page request.PageFilter,
) (*response.PageResponse[*models.SysRole], error) {
	return c.roleSvc.GetByFilter(ctx, filter, page)
}
func (c *roleControllerImpl) Tree(ctx *gin.Context) ([]*dto.RoleTree, error) {
	return c.roleSvc.Tree(ctx)
}

func (c *roleControllerImpl) Create(ctx *gin.Context, req dto.RoleRequestForm) error {
	_, err := c.roleSvc.Create(ctx, req)
	return err
}
func (c *roleControllerImpl) UpdateByID(ctx *gin.Context, id int, req dto.RoleRequestForm) error {
	_, err := c.roleSvc.UpdateByID(ctx, uint64(id), req)
	return err
}
func (c *roleControllerImpl) Delete(ctx *gin.Context, id int) error {
	return c.roleSvc.DeleteByID(ctx, uint64(id))
}
