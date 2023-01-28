package container

import (
	"app/container"
	"app/modules/system/controller"
	"app/modules/system/dao"
	"app/modules/system/routes"
	"app/modules/system/service"
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
