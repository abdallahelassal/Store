package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DB_NAME		string
	DB_PASSWORD	string
	DB_USER		string
	DB_PORT		string
	DB_HOST		string
	SECRET_KEY	string
	PORT 		string
}

var AppConfig Config

func LoadConfig(fileName string){
	err := godotenv.Load(fileName)
	if err != nil {
		log.Printf("error log .env file %v", err)
		return
	}
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
		return
	}

	AppConfig = Config{
		DB_NAME: os.Getenv("DB_NAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_USER: os.Getenv("DB_USER"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_HOST: os.Getenv("DB_HOST"),
		SECRET_KEY: os.Getenv("SECRET_KEY"),
		PORT: port,
	}
}