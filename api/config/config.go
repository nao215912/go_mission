package config

import (
	"fmt"
	"os"
)

const (
	portKey     = "PORT"
	defaultPort = "8080"
)

func Port() string {
	port, err := getString(portKey)
	if err != nil {
		return defaultPort
	}
	return port
}

func getString(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return "", fmt.Errorf("config:[%s] not found", key)
	}
	return v, nil
}
