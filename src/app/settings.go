package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type settings struct {
	HomeDir string
	Domain  string
	Host    string
	Port    string

	DbHost     string
	DbPort     int64
	DbName     string
	DbUser     string
	DbPassword string
}

var Settings = new(settings)

func ReadSettings() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found or could not load it: %v", err)
	}

	Settings.HomeDir = os.Getenv("HOME_DIR")
	Settings.Domain = os.Getenv("DOMAIN")
	Settings.Host = os.Getenv("HOST")
	Settings.Port = os.Getenv("PORT")

	Settings.DbHost = os.Getenv("DB_HOST")
	Settings.DbName = os.Getenv("DB_NAME")
	Settings.DbUser = os.Getenv("DB_USER")
	Settings.DbPassword = os.Getenv("DB_PASSWORD")
	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr != "" {
		dbPort, err := strconv.ParseInt(dbPortStr, 10, 64)
		if err != nil {
			log.Printf("Invalid DB_PORT value (%s). Defaulting to 5432", dbPortStr)
			dbPort = 5432
		}
		Settings.DbPort = dbPort
	} else {
		Settings.DbPort = 1433
	}
	fmt.Println("Settings read from .env file")
}
