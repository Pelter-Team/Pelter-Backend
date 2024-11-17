package utils

import (
	"strconv"

	"Pelter_backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func ParseIDParam(ctx *fiber.Ctx) (uint, error) {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return 0, ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error:   "Invalid ID",
			Success: false,
		})
	}
	return uint(id), nil
}
