package product

import (
	"Pelter_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Route(app *fiber.App, gorm *gorm.DB) {
	repo := NewProductRepository(gorm)
	usecase := NewProductUsecase(repo)
	service := NewProductService(usecase)

	group := app.Group("/products")
	group.Get("/", middleware.ValidateCookie, service.InsertProduct)
}
