package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	GRPCHost string
	GRPCPort string
	HTTPPort string
}

func Load() *Config {
	grpcHost := os.Getenv("GRPC_HOST")
	if grpcHost == "" {
		grpcHost = "localhost"
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	log.Printf("Config loaded: GRPC_HOST=%s, HTTP_PORT=%s", grpcHost, httpPort)

	return &Config{
		GRPCHost: grpcHost,
		GRPCPort: grpcPort,
		HTTPPort: httpPort,
	}
}

func (c *Config) GRPCAddress() string {
	return fmt.Sprintf("%s:%s", c.GRPCHost, c.GRPCPort)
}
