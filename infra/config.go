package infra

import (
	"os"
)

type Config struct {
	DBName string
}

var config *Config

func LoadConfig() *Config {
	config = &Config{
		DBName: getEnv("DB_NAME"),
	}

	return config
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic("Chave não encontrada nas variavis de ambiente: " + key)
}
