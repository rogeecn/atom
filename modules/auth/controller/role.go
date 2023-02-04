package controller

import (
	"atom/modules/auth/dto"
	"atom/providers/config"

	"github.com/gin-gonic/gin"
)

type RoleController interface {
	GetName(*gin.Context) (string, error)
}

type roleControllerImpl struct {
	conf *config.Config
}

func NewRoleController(conf *config.Config) RoleController {
	return &roleControllerImpl{conf: conf}
}

func (c *roleControllerImpl) GetName(ctx *gin.Context) (string, error) {
	return "Role", nil
}

func (c *roleControllerImpl) Create(ctx *gin.Context, req *dto.RoleCreateRequest) error {
	return nil
}
