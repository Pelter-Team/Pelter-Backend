package main

import (
	"Pelter_backend/internal/config"
	"Pelter_backend/internal/repository"
)

func main() {

	cfg := config.LoadConfig()

	_ = cfg
	repo := repository.NewRepository()
	repo.InsertProduct()

}
