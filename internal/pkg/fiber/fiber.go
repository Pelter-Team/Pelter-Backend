package fiber

import (
	"Pelter_backend/internal/config"

	"github.com/gofiber/fiber/v2"
)

func FiberConn(cfg *config.App) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: cfg.Name,
	})
	return app
}
