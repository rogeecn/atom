package controller

import (
	"atom/providers/jwt"
	"atom/providers/rbac"

	"github.com/gin-gonic/gin"
)

type PermissionController interface {
	Get(ctx *gin.Context) (string, error)
}

type permissionControllerImpl struct {
	jwt  *jwt.JWT
	rbac rbac.IRbac
}

func NewPermissionController(
	jwt *jwt.JWT,
	rbac rbac.IRbac,
) PermissionController {
	return &permissionControllerImpl{rbac: rbac, jwt: jwt}
}

func (c *permissionControllerImpl) Get(ctx *gin.Context) (string, error) {
	claims, err := c.jwt.GetClaims(ctx)
	if err != nil {
		return "", err
	}

	perm, err := c.rbac.JsonPermissionsForUser(claims.Username)
	if err != nil {
		return "", err
	}

	return perm, nil
}
