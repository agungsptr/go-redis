package db

import (
	"fmt"

	"github.com/agungsptr/go-redis/config"
	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Get().RedisHost, config.Get().RedisPort),
		Username: config.Get().RedisUser,
		Password: config.Get().RedisPass,
		DB:       0,
	})
	return client
}
