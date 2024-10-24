package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Route(app *fiber.App, gorm *gorm.DB) {
	repo := NewProductRepository(gorm)
	usecase := NewProductUsecase(repo)
	service := NewProductService(usecase)

	group := app.Group("/products")
	group.Get("/", service.InsertProduct)
}
