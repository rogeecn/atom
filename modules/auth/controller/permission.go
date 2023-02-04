package controller

import (
	"atom/providers/config"

	"github.com/gin-gonic/gin"
)

type PermissionController interface {
	GetName(*gin.Context) (string, error)
}

type permissionControllerImpl struct {
	conf *config.Config
}

func NewPermissionController(conf *config.Config) PermissionController {
	return &permissionControllerImpl{conf: conf}
}

func (c *permissionControllerImpl) GetName(ctx *gin.Context) (string, error) {
	return "Permission",nil
}
