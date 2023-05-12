package services

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers/grpcs"
	"go.uber.org/dig"
)

type GrpcService struct {
	dig.In

	Server   *grpcs.Grpc
	Services []grpcs.ServerService `group:"grpc_server_services"`
}

func ServeGrpc() error {
	return container.Container.Invoke(func(grpc GrpcService) error {
		for _, svc := range grpc.Services {
			grpc.Server.RegisterService(svc.Name(), svc.Register)
		}
		return grpc.Server.Serve()
	})
}