package routes

import (
	"atom/contracts"
	"atom/modules/auth/controller"
	"atom/providers/http"

	"github.com/rogeecn/gen"
)

type Route struct {
	svc        *http.Service
	role       controller.RoleController
	permission controller.PermissionController
}

func NewRoute(
	svc *http.Service,
	role controller.RoleController,
	permission controller.PermissionController,
) contracts.Route {
	return &Route{
		svc:        svc,
		role:       role,
		permission: permission,
	}
}

func (r *Route) Register() {
	group := r.svc.Engine.Group("auth")
	{
		roleGroup := group.Group("role")
		{
			roleGroup.GET("/roles", gen.DataFunc(r.role.GetName))
		}

		permissionGroup := group.Group("permission")
		{
			permissionGroup.GET("/permissions", gen.DataFunc(r.permission.GetName))
		}

	}
}
