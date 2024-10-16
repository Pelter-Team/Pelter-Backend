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

func (r *productService) InsertProduct(ctx *fiber.Ctx) error {
	_ = r.productUsecase.InsertProduct(ctx.Context())

	return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
		Error: "Bad Request",
	})
}
