package contracts

type HttpRoute interface{}

type HttpService interface {
	Serve() error
	GetEngine() interface{}
}
