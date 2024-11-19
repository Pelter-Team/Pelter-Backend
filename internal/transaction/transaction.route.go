package transaction

import (
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
	groups.Get("/", service.GetTransactions)
	group := app.Group("/transaction")
	group.Post("/buy/:id", service.CreateTransaction)
	group.Get("/:id", service.GetTransactionByID)
	group.Get("/user/:id", service.GetTransactionsByUserID)
}
