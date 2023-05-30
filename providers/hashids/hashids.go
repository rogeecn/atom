package hashids

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers/log"
	"github.com/rogeecn/atom/utils/opt"

	"github.com/speps/go-hashids/v2"
)

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Config
	if err := o.UnmarshalConfig(&config); err != nil {
		log.Fatal(err)
	}
	return container.Container.Provide(func() (*hashids.HashID, error) {
		return hashids.NewWithData(&hashids.HashIDData{
			MinLength: int(config.MinLength),
			Salt:      config.Salt,
		})
	}, o.DiOptions()...)
}
