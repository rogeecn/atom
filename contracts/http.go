package contracts

type Route interface{}

type HttpService interface {
	Serve() error
	GetEngine() interface{}
}
