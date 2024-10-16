package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Route(app *fiber.App, gorm *gorm.DB) {
	repo := NewProductRepository(gorm)
	usecase := NewProductUsecase(repo)
	service := NewProductService(usecase)

	group := app.Group("/product_v1")
	group.Post("/insert-product", service.InsertProduct)
}
