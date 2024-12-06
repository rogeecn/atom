package contracts

type HttpRoute interface {
	Register() any
	Name() string
}
