package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/caarlos0/env/v11"
)

type configurations struct {
	RedisHost string `env:"REDIS_HOST"`
	RedisPort string `env:"REDIS_PORT"`
	RedisUser string `env:"REDIS_USER"`
	RedisPass string `env:"REDIS_PASS"`
	MongoUser string `env:"MONGO_USER"`
	MongoPass string `env:"MONGO_PASS"`
	MongoPort string `env:"MONGO_PORT"`
	MongoHost string `env:"MONGO_HOST"`
	MongoUri  string `env:"omitempty"`
}

var (
	cfg  configurations
	once sync.Once
)

func loadEnv() {
	if err := env.Parse(&cfg); err != nil {
		log.Panic(err.Error())
	}

	cfg.MongoUri = fmt.Sprintf("mongodb://%s:%s@%s:%s",
		cfg.MongoUser,
		cfg.MongoPass,
		cfg.MongoHost,
		cfg.MongoPort,
	)
}

func Get() configurations {
	once.Do(loadEnv)
	return cfg
}
