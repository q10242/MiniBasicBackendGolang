package base

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort   string
	DBHost    string
	DBPort    string
	DBUser    string
	DBName    string
	DBPass    string
	JWTSecret string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // 支援環境變數覆蓋

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = Config{
		AppPort:   viper.GetString("APP_PORT"),
		DBHost:    viper.GetString("DB_HOST"),
		DBPort:    viper.GetString("DB_PORT"),
		DBUser:    viper.GetString("DB_USER"),
		DBName:    viper.GetString("DB_NAME"),
		DBPass:    viper.GetString("DB_PASS"),
		JWTSecret: viper.GetString("JWT_SECRET"),
	}
}
