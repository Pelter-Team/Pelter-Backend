package server

import (
	"Pelter_backend/internal/config"
	"Pelter_backend/internal/product"
	"Pelter_backend/internal/user"
	"context"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type (
	server struct {
		app *fiber.App
		cfg *config.App
	}
)

func Start(pctx context.Context, cfg *config.App, app *fiber.App, gorm *gorm.DB) {
	s := &server{
		cfg: cfg,
		app: app,
	}

	product.Route(s.app, gorm)
	user.Route(s.app, gorm)
}
