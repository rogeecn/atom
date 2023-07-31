package contracts

type MicroService interface {
	Serve() error
	GetEngine() any
}
