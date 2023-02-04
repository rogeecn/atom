package container

import (
	"atom/container"
	"atom/modules/auth/controller"
	"atom/modules/auth/dao"
	"atom/modules/auth/routes"
	"log"

	"go.uber.org/dig"
)

func init() {

	// controller
	if err := container.Container.Provide(controller.NewRoleController); err != nil {
		log.Fatal(err)
	}

	if err := container.Container.Provide(controller.NewPermissionController); err != nil {
		log.Fatal(err)
	}

	//service

	// dao
	if err := container.Container.Provide(dao.NewRoleDao); err != nil {
		log.Fatal(err)
	}

	if err := container.Container.Provide(routes.NewRoute, dig.Group("route")); err != nil {
		log.Fatal(err)
	}
}
