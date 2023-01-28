package micro

import (
	"atom/container"
	"atom/providers/config"
	"atom/utils"
	"log"

	mgrpc "github.com/go-micro/plugins/v4/client/grpc"
	mhttp "github.com/go-micro/plugins/v4/server/http"
	"go-micro.dev/v4"
)

func init() {
	if err := container.Container.Provide(NewService); err != nil {
		log.Fatal(err)
	}
}

func NewService(conf *config.Config) micro.Service {
	service := micro.NewService(
		micro.Server(mhttp.NewServer()),
		micro.Client(mgrpc.NewClient()),
		micro.Address(conf.Http.Address()),
	)
	service.Init(micro.Name(utils.Service), micro.Version(utils.Version))

	return service
}
