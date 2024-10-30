package product

import (
	"Pelter_backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

type (
	productService struct {
		productUsecase ProductUsecase
	}
	ProductService interface {
		InsertProduct(ctx *fiber.Ctx) error
	}
)

func NewProductService(productUsecase ProductUsecase) ProductService {
	return &productService{
		productUsecase: productUsecase,
	}
}

func (s *productService) InsertProduct(ctx *fiber.Ctx) error {
	_ = s.productUsecase.InsertProduct(ctx.Context())

	return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
		Error: "Authenticated",
	})
}
