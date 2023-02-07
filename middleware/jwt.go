package middleware

import (
	"atom/providers/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogeecn/gen"
)

func JWTAuth(j *jwt.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			gen.NewBusError(http.StatusBadRequest, http.StatusBadRequest, "未登录或非法访问").JSON(c, false)
			c.Abort()
			return
		}

		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			gen.NewBusError(http.StatusBadRequest, http.StatusBadRequest, err.Error()).JSON(c, false)
			c.Abort()
			return
		}

		// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 如果需要 请自行打开
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

		//if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}

		c.Set(jwt.CtxKey, claims)
		c.Next()
	}
}
