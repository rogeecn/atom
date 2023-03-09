package middleware

import (
	"atom/providers/config"
	"atom/providers/jwt"
	"atom/providers/rbac"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rogeecn/gen"
)

// Permission 拦截器
func CheckPermission(config *config.Config, rbac rbac.IRbac) gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.App.Mode != "production" {
			c.Next()
			return
		}

		claimsCtx, exists := c.Get(jwt.CtxKey)
		if !exists {
			gen.NewBusError(http.StatusBadRequest, http.StatusBadRequest, "Token 获取失败").JSON(c, false)
			c.Abort()
			return
		}
		claims := claimsCtx.(jwt.Claims)

		// 获取请求的PATH
		path := c.Request.URL.Path

		// 获取请求方法
		method := c.Request.Method

		// 获取用户的角色
		role := strconv.Itoa(int(claims.Role))

		if !rbac.Can(role, method, path) {
			gen.NewBusError(http.StatusForbidden, http.StatusForbidden, "未登录或非法访问").JSON(c, false)
			c.Abort()
			return
		}
	}
}
