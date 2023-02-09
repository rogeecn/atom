package controller

import (
	"atom/common/request"
	"atom/common/response"
	"atom/database/models"
	"atom/modules/auth/dto"
	"atom/modules/auth/service"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleSvc service.RoleService
}

func NewRoleController(
	roleSvc service.RoleService,
) *RoleController {
	return &RoleController{
		roleSvc: roleSvc,
	}
}

func (c *RoleController) GetByFilter(
	ctx *gin.Context,
	filter dto.RoleRequestFilter,
	page request.PageFilter,
) (*response.PageResponse[*models.SysRole], error) {
	return c.roleSvc.GetByFilter(ctx, filter, page)
}
func (c *RoleController) Tree(ctx *gin.Context) ([]*dto.RoleTree, error) {
	return c.roleSvc.Tree(ctx)
}

func (c *RoleController) Create(ctx *gin.Context, req dto.RoleRequestForm) error {
	_, err := c.roleSvc.Create(ctx, req)
	return err
}
func (c *RoleController) UpdateByID(ctx *gin.Context, id int, req dto.RoleRequestForm) error {
	_, err := c.roleSvc.UpdateByID(ctx, uint64(id), req)
	return err
}
func (c *RoleController) Delete(ctx *gin.Context, id int) error {
	return c.roleSvc.DeleteByID(ctx, uint64(id))
}
