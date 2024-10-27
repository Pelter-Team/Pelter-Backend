package middleware

import (
	"Pelter_backend/internal/dto"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error { // validate the struct using the validator package
	return validate.Struct(s)
}

func ValidationMiddleware(s interface{}) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := ctx.BodyParser(s); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
				Error: "Invalid request format",
			})
		}

		if err := ValidateStruct(s); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
				Error: err.Error(),
			})
		}
		ctx.Locals("body", s) // Store the validated struct in the context for later use

		return ctx.Next()
	}
}
