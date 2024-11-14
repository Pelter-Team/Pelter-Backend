package transaction

import (
	"Pelter_backend/internal/dto"
	"Pelter_backend/internal/entity"
	"fmt"

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
		ListAllTransactionsByUserId(ctx *fiber.Ctx) error
	}
)

func NewTransactionService(transactionUsecase TransactionUsecase) TransactionService {
	return &transactionService{
		transactionUsecase: transactionUsecase,
	}
}

func (s *transactionService) Buy(ctx *fiber.Ctx) error {
	var req dto.TransactionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}
	txn := entity.Transaction{
		ProductID: req.ProductID,
		BuyerID:   req.BuyerID,
	}
	if err := s.transactionUsecase.Buy(ctx.UserContext(), &txn); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.HttpResponse{
		Result:  "Transaction success",
		Success: true,
	})
}

func (s *transactionService) FindById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	txn, err := s.transactionUsecase.FindById(ctx.UserContext(), uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(txn)
}

func (s *transactionService) ListAllTransactions(ctx *fiber.Ctx) error {
	txns, err := s.transactionUsecase.ListAll(ctx.UserContext())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(txns)
}

func (s *transactionService) ListAllTransactionsByUserId(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	txns, err := s.transactionUsecase.ListAllByUserId(ctx.UserContext(), uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(txns)
}
