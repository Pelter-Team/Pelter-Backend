package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParseIDParam(ctx *fiber.Ctx) (uint, error) {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
