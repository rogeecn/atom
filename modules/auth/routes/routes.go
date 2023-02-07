package routes

import (
	"atom/common/err"
	"atom/common/request"
	"atom/contracts"
	"atom/modules/auth/controller"
	"atom/modules/auth/dto"
	"atom/providers/http"

	"github.com/rogeecn/gen"
)

type Route struct {
	svc        *http.Service
	user       controller.UserController
	role       controller.RoleController
	permission controller.PermissionController
}

func NewRoute(
	svc *http.Service,
	role controller.RoleController,
	user controller.UserController,
	permission controller.PermissionController,
) contracts.Route {
	return &Route{
		svc:        svc,
		role:       role,
		user:       user,
		permission: permission,
	}
}

func (r *Route) Register() {
	group := r.svc.Engine.Group("auth")
	{
		// common functions
		group.GET("/login", gen.DataFunc1(
			r.user.Login,
			gen.BindBody(dto.LoginRequestForm{}, err.BindBodyFailed),
		))

		roleGroup := group.Group("roles")
		{
			roleGroup.GET("", gen.DataFunc2(
				r.role.GetByFilter,
				gen.BindQuery(dto.RoleRequestFilter{}, err.BindQueryFailed),
				gen.BindQuery(request.PageFilter{}, err.BindQueryFailed),
			))

			roleGroup.GET("/tree", gen.DataFunc(r.role.Tree))

			roleGroup.POST("", gen.Func1(
				r.role.Create,
				gen.BindBody(dto.RoleRequestForm{}, err.BindBodyFailed),
			))

			roleGroup.PUT("/:id", gen.Func2(
				r.role.UpdateByID,
				gen.Int("role_id", err.BindPathFailed.Format("id")),
				gen.BindBody(dto.RoleRequestForm{}, err.BindBodyFailed),
			))

			roleGroup.DELETE("/:id", gen.Func1(
				r.role.Delete,
				gen.Int("role_id", err.BindPathFailed.Format("id")),
			))
		}

		permissionGroup := group.Group("permission")
		{
			permissionGroup.GET("/permissions", gen.DataFunc(r.permission.Get))
		}

	}
}
