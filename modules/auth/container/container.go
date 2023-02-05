package container

import (
	"atom/container"
	"atom/modules/auth/controller"
	"atom/modules/auth/dao"
	"atom/modules/auth/routes"
	"atom/modules/auth/service"
	"log"

	"go.uber.org/dig"
)

func init() {

	// controller
	if err := container.Container.Provide(controller.NewRoleController); err != nil {
		log.Fatal(err)
	}

	if err := container.Container.Provide(controller.NewUserController); err != nil {
		log.Fatal(err)
	}

	if err := container.Container.Provide(controller.NewPermissionController); err != nil {
		log.Fatal(err)
	}

	//service
	if err := container.Container.Provide(service.NewRoleService); err != nil {
		log.Fatal(err)
	}

	if err := container.Container.Provide(service.NewUserService); err != nil {
		log.Fatal(err)
	}

	// dao
	if err := container.Container.Provide(dao.NewRoleDao); err != nil {
		log.Fatal(err)
	}

	if err := container.Container.Provide(dao.NewUserDao); err != nil {
		log.Fatal(err)
	}

	if err := container.Container.Provide(dao.NewUserRoleDao); err != nil {
		log.Fatal(err)
	}

	// routes
	if err := container.Container.Provide(routes.NewRoute, dig.Group("route")); err != nil {
		log.Fatal(err)
	}
}
