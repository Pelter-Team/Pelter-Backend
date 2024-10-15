package product

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	// TODO: send gorm db here
	repo := NewProductRepository(nil)
	usecase := NewProductUsecase(repo)
	service := NewProductService(usecase)

	group := app.Group("/product_v1")
	group.Post("/insert-product", service.InsertProduct)
}
