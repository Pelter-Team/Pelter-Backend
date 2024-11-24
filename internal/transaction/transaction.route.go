package transaction

import (
	"Pelter_backend/internal/middleware"
	"Pelter_backend/internal/product"
	"Pelter_backend/internal/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Route(app *fiber.App, gorm *gorm.DB) {
	transactionRepo := NewTransactionRepository(gorm)
	productRepo := product.NewProductRepository(gorm)
	userRepo := user.NewUserRepository(gorm)
	usecase := NewTransactionUsecase(transactionRepo, userRepo, productRepo)
	service := NewTransactionService(usecase)

	groups := app.Group("/transactions")

	groups.Get("/", middleware.ValidateCookie, service.GetTransactions)
	group := app.Group("/transaction")
	group.Post("/buy/:id", middleware.ValidateCookie, service.CreateTransaction)
	group.Get("/:id", middleware.ValidateCookie, service.GetTransactionByID)
	group.Get("/user/:id", middleware.ValidateCookie, service.GetTransactionsByUserID)
}
