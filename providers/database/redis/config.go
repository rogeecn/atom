package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

const DefaultPrefix = "Redis"

type Config struct {
	Host     string
	Port     uint
	Password string
	DB       uint
}

func (c *Config) ToRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       int(c.DB),
	}
}
