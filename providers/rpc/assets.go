package rpc

import (
	"log"

	"atom/container"
	"atom/proto"

	"go-micro.dev/v4"
)

func init() {
	err := container.Container.Provide(func(svc micro.Service) proto.WebApiService {
		return proto.NewWebApiService("web.api", svc.Client())
	})

	if err != nil {
		log.Fatal(err)
	}
}
