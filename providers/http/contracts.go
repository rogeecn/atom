package http

type Route interface {
	Register()
}

type Service interface {
	Serve() error
	GetEngine() interface{}
}
