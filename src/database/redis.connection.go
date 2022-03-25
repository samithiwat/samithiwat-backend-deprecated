package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/samithiwat/samithiwat-backend/src/config"
)

var ctx = context.Background()

type Cache interface {
	getConnection() *redis.Client
}

type cache struct {
	config     *config.Config
	connection *redis.Client
}

func InitRedisConnect(config *config.Config) (Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       0,
	})

	return &cache{connection: client, config: config}, nil
}

func (c *cache) getConnection() *redis.Client {
	return c.connection
}
