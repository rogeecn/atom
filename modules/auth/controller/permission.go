package controller

import (
	"atom/providers/jwt"
	"atom/providers/log"
	"atom/providers/rbac"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogeecn/gen"
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
	claimsCtx, exists := ctx.Get(jwt.CtxKey)
	if !exists {
		return "", gen.NewBusError(http.StatusBadRequest, http.StatusBadRequest, "Token 获取失败")
	}
	claims := claimsCtx.(jwt.Claims)
	log.Debug("claim: ", claims)

	perm, err := c.rbac.JsonPermissionsForUser("Rogee")
	if err != nil {
		return "", err
	}

	return perm, nil
}
