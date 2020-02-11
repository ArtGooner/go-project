package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	ConnectionString string `env:"CONNECTION_STRING" envDefault:"host=localhost port=5432 user=postgres password=213612458 dbname=therion sslmode=disable"`
	HttpPort         int    `env:"HTTP_PORT" envDefault:"8888"`
	GrpcPort         string `env:"GRPC_PORT" envDefault:":50051"`
	GrpcAddress      string `env:"GRPC_ADDRESS" envDefault:"localhost:50051"`
}

// New config from environment variables
func New() (*Config, error) {
	c := &Config{}
	err := env.Parse(c)
	return c, err
}
