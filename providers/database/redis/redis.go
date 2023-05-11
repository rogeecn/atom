package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var conf Config
	if err := o.UnmarshalConfig(&conf); err != nil {
		return err
	}

	return container.Container.Provide(func() (*redis.Client, error) {
		client := redis.NewClient(conf.ToRedisOptions())

		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			return nil, errors.Wrap(err, "failed to ping")
		}

		return client, nil
	}, o.DiOptions()...)
}
