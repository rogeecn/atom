package container

import (
	"atom/container"
	"atom/modules/system/controller"
	"atom/modules/system/dao"
	"atom/modules/system/routes"
	"atom/modules/system/service"
	"log"

	"go.uber.org/dig"
)

func init() {
	if err := container.Container.Provide(dao.NewDao); err != nil {
		log.Fatal(err)
	}

	if err := container.Container.Provide(service.NewSystemService); err != nil {
		log.Fatal(err)
	}

	if err := container.Container.Provide(controller.NewController); err != nil {
		log.Fatal(err)
	}

	if err := container.Container.Provide(routes.NewRoute, dig.Group("route")); err != nil {
		log.Fatal(err)
	}
}
