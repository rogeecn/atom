package controller

import (
	"atom/modules/auth/dto"
	"atom/modules/auth/service"
	"atom/providers/config"
	"atom/providers/jwt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	conf *config.Config
	user service.UserService
	jwt  *jwt.JWT
}

func NewUserController(
	conf *config.Config,
	user service.UserService,
	jwt *jwt.JWT,
) *UserController {
	return &UserController{
		conf: conf,
		user: user,
		jwt:  jwt,
	}
}

func (c *UserController) Login(ctx *gin.Context, req dto.LoginRequestForm) (*dto.LoginResponse, error) {
	user, err := c.user.AuthMatchPassword(ctx, &req)
	if err != nil {
		return nil, err
	}

	token, err := c.user.GenerateJWTTokenFromUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
	}, nil
}
