package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App        App
		Database   Db
		Cloudinary Cloudinary
	}

	App struct {
		Port   string
		Name   string
		Stage  string
		Domain string
		Origin string
	}

	Db struct {
		Url string
	}

	Cloudinary struct {
		Cloudname string
		Apikey    string
		Apisecret string
	}
)

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	cfg := &Config{
		App: App{
			Port: os.Getenv("PORT"),
			Name: os.Getenv("APP_NAME"),
		},
		Database: Db{
			Url: os.Getenv("DATABASEURL"),
		},
		Cloudinary: Cloudinary{
			Cloudname: os.Getenv("CLOUDNAME"),
			Apikey:    os.Getenv("APIKEY"),
			Apisecret: os.Getenv("APISECRET"),
		},
	}
	return cfg
}

//func (d *Database) DbConn() {
//	fmt.Println(d.DatabaseUrl)
//}
