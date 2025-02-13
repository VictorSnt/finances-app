package infra

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName string
}

var config *Config

func LoadConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Arquivo .env não encontrado. Usando variáveis de ambiente do sistema.")
	}

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
