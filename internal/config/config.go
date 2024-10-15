package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App      App
	Database Database
}

type App struct {
	Port string
}

type Database struct {
	DatabaseUrl string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	databaseUrl := os.Getenv("DATABASEURL")
	port := os.Getenv("PORT")

	cfg := &Config{
		App: App{
			Port: port,
		},
		Database: Database{
			DatabaseUrl: databaseUrl,
		},
	}
	return cfg
}

//func (d *Database) DbConn() {
//	fmt.Println(d.DatabaseUrl)
//}
