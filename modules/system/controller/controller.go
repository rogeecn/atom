package controller

import (
	"app/modules/system/dao"
	"app/modules/system/dto"
	"app/modules/system/service"
	"app/providers/config"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetName(*gin.Context) (dto.Name, error)
}

type ControllerImpl struct {
	Conf *config.Config
	svc  service.SystemService
}

func NewController(Conf *config.Config, dao dao.Dao, svc service.SystemService) Controller {
	return &ControllerImpl{Conf: Conf, svc: svc}
}

func (c *ControllerImpl) GetName(ctx *gin.Context) (dto.Name, error) {
	return c.svc.GetName(ctx)
}
