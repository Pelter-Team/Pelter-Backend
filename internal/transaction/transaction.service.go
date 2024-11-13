package transaction

import (
	"github.com/gofiber/fiber/v2"
)

type (
	transactionService struct {
		transactionUsecase TransactionUsecase
	}
	TransactionService interface {
		Buy(ctx *fiber.Ctx) error
		FindById(ctx *fiber.Ctx) error
		ListAllTransactions(ctx *fiber.Ctx) error
	}
)

func NewTransactionService(transactionUsecase TransactionUsecase) TransactionService {
	return &transactionService{
		transactionUsecase: transactionUsecase,
	}
}

func (s *transactionService) Buy(ctx *fiber.Ctx) error {
	var req dto
