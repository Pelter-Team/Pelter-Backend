package transaction

import (
	"Pelter_backend/internal/dto"
	"Pelter_backend/internal/entity"
	"Pelter_backend/internal/pkg/jwt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type (
	transactionService struct {
		transactionUsecase TransactionUsecase
	}
	TransactionService interface {
		CreateTransaction(ctx *fiber.Ctx) error
		GetTransactions(ctx *fiber.Ctx) error
		GetTransactionByID(ctx *fiber.Ctx) error
		GetTransactionsByUserID(ctx *fiber.Ctx) error
	}
)

func NewTransactionService(transactionUsecase TransactionUsecase) TransactionService {
	return &transactionService{
		transactionUsecase: transactionUsecase,
	}
}

func (s *transactionService) CreateTransaction(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error:   "Invalid ID",
			Success: false,
		})
	}
	userId, err := jwt.GetIDFromToken(ctx.Cookies("access_token"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   "Cannot get UserID from access_token context: " + err.Error(),
			Success: false,
		})
	}
	txn := entity.Transaction{
		ProductID: uint(id),
		BuyerID:   userId,
	}
	if err := s.transactionUsecase.CreateTransaction(ctx.UserContext(), &txn); err != nil {
		if err.Error() == "Product already sold" {
			return ctx.Status(fiber.StatusConflict).JSON(dto.HttpResponse{
				Error:   err.Error(),
				Success: false,
			})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error:   err.Error(),
			Success: false,
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.HttpResponse{
		Result: dto.TransactionResponse{
			ID:        txn.ID,
			ProductID: txn.ProductID,
			BuyerID:   txn.BuyerID,
			Amount:    uint(txn.Amount),
			CreatedAt: txn.CreatedAt.String(),
		},
		Success: true,
	})
}

func (s *transactionService) GetTransactions(ctx *fiber.Ctx) error {
	txns, err := s.transactionUsecase.GetTransactions(ctx.UserContext())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(txns)
}

func (s *transactionService) GetTransactionByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error:   "Invalid ID",
			Success: false,
		})
	}

	txn, err := s.transactionUsecase.GetTransactionByID(ctx.UserContext(), uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(txn)
}

func (s *transactionService) GetTransactionsByUserID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error:   "Invalid ID",
			Success: false,
		})
	}
	txns, err := s.transactionUsecase.GetTransactionsByUserID(ctx.UserContext(), uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(txns)
}