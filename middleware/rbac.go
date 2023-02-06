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
func CheckPermission(config *config.Config, rbac rbac.IRbac, jwt *jwt.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.App.Mode != "production" {
			c.Next()
			return
		}

		claim, err := jwt.GetClaims(c)
		if err != nil {
			gen.NewBusError(http.StatusBadRequest, http.StatusBadRequest, "Token 获取失败").JSON(c, false)
			c.Abort()
			return
		}

		//获取请求的PATH
		path := c.Request.URL.Path

		// 获取请求方法
		method := c.Request.Method

		// 获取用户的角色
		role := strconv.Itoa(int(claim.RoleID))

		if rbac.Can(role, method, path) == false {
			gen.NewBusError(http.StatusForbidden, http.StatusForbidden, "未登录或非法访问").JSON(c, false)
			c.Abort()
			return
		}
	}
}
