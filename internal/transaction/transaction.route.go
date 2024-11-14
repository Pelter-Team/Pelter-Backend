package transaction

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Route(app *fiber.App, gorm *gorm.DB) {
	repo := NewTransactionRepository(gorm)
	usecase := NewTransactionUsecase(repo)
	service := NewTransactionService(usecase)

	route := app.Group("/transactions")
	route.Post("/buy", service.Buy)
	route.Get("/:id", service.FindById)
	route.Get("/", service.ListAllTransactions)
	route.Get("/user/:id", service.ListAllTransactionsByUserId)
}
