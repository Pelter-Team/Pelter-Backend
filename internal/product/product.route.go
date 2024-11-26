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

	//can be access by anyone
	groups := app.Group("/products")
	// TODO: this one doesn't implement correctly in dto so it return all the field
	// TODO: maybe it will be unused
	// groups.Get("/buyer", middleware.ValidateCookie, service.GetProductByBuyerId)
	groups.Get("/user", middleware.ValidateCookie, service.GetProductByUserId)

	// be post for now since need an array of productId
	groups.Post("/favorite", service.GetProductIn)

	groups.Get("/", service.GetProduct)
	group := app.Group("/product")
	group.Get(("/:id"), service.GetProductByID)

	//below = secure (need to have a valid cookie)
	group.Post("/add", middleware.ValidateCookie, service.InsertProduct)
	group.Put(("/:id"), middleware.ValidateCookie, service.UpdateProduct)
	group.Delete(("/:id"), middleware.ValidateCookie, service.DeleteProduct)
	group.Patch("/verification/:id", service.UpdateProductVerificationStatus)
	group.Patch("/is-sold/:id", middleware.ValidateCookie, service.UpdateProductIsSold)

	//below = admin only
	adminOnlyGroup := group.Group("/admin")
	adminOnlyGroup.Put("/:id", middleware.ValidateCookie, service.UpdateProductAdmin)
	adminOnlyGroup.Delete(("/:id"), middleware.ValidateCookie, service.DeleteProductAdmin)
}
