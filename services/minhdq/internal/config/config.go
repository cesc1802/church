package config

import (
	"github.com/kelseyhightower/envconfig"
	"sync"
)

var (
	config    Config
	mu        sync.RWMutex
	envPrefix = "CHURCH"
)

type Config struct {
	PostgresURL string `envconfig:"postgres_url"`
	HttpURL     string `envconfig:"http_url"`
	HttpPort    int    `evnconfig:"http_port"`
}

func Get() Config {
	mu.RLock()
	defer mu.RUnlock()
	return config
}

func Set(c Config) {
	mu.Lock()
	defer mu.Unlock()
	config = c
}

func LoadFromEnv() (err error) {
	mu.Lock()
	defer mu.Unlock()
	return envconfig.Process(envPrefix, &config)
}
