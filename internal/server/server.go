package server

import (
	"Pelter_backend/internal/config"
	"Pelter_backend/internal/product"
	"context"

	"github.com/gofiber/fiber/v2"
)

type (
	server struct {
		app *fiber.App
		cfg *config.App
	}
)

func Start(pctx context.Context, cfg *config.App, app *fiber.App) {
	s := &server{
		cfg: cfg,
		app: app,
	}
	_ = s.cfg

	product.Route(s.app)

}
