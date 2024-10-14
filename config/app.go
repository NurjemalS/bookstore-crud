package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {

	DbHost       string `mapstructure:"db_host"`
	DbPort       string `mapstructure:"db_port"`
	DbDatabase   string `mapstructure:"db_database"`
	DbUsername   string `mapstructure:"db_username"`
	DbPassword   string `mapstructure:"db_password"`
	ServerAddress string `mapstructure:"server_address"`
}

func LoadConfig() *Config {

	err := godotenv.Load()
    if err != nil {
 	   log.Fatalf("Error loading .env file")
    }
	
	return &Config{

		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbDatabase:  os.Getenv("DB_NAME"),
		DbUsername:  os.Getenv("DB_USER"),
		DbPassword:  os.Getenv("DB_PASSWORD"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
	}
}