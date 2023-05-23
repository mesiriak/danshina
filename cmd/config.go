package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"log"
)

type Config struct {
	DebugMode bool `env:"DEBUG_MODE"`

	AppHost string `env:"APP_HOST"`
	AppPort int    `env:"APP_PORT"`

	MongoUrl string `env:"MONGO_URL"`
}

var config *Config

func GetConfig() Config {
	if config != nil {
		return *config
	}

	config = &Config{}

	err := env.Parse(config)

	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
	// FIXME: config doesn't load
	fmt.Println(config)
	return *config
}
