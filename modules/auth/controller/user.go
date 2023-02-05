package controller

import (
	"atom/providers/config"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetName(*gin.Context) (string, error)
}

type userControllerImpl struct {
	conf *config.Config
}

func NewUserController(conf *config.Config) UserController {
	return &userControllerImpl{conf: conf}
}

func (c *userControllerImpl) GetName(ctx *gin.Context) (string, error) {
	return "User",nil
}
