package middleware

import (
	"Pelter_backend/internal/dto"
	"Pelter_backend/internal/pkg/jwt"

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
				Error: "Invalid request format " + err.Error(),
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

func ValidateCookie(ctx *fiber.Ctx) error {
	cookieValue := ctx.Cookies("access_token")

	if cookieValue == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
			Error: "Unauthorized: missing cookie",
		})
	}

	_, err := jwt.ValidateToken(cookieValue)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
			Error: "Unauthorized: " + err.Error(),
		})
	}
	// ctx.Locals("claims", claims)
	return ctx.Next()
}

// func ValidateRole(ctx *fiber.Ctx) error {
// 	claims, err := jwt.ValidateToken(ctx.Cookies("access_token"))
// 	if err != nil {
// 		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
// 			Error: "Unauthorized: " + err.Error(),
// 		})
// 	}

// 	role := claims["role"].(string)
// 	if role != "admin" {
// 		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
// 			Error: "Unauthorized: invalid role",
// 		})
// 	}
// 	return ctx.Next()
// }
