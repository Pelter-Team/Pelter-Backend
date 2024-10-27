package user

import (
	"Pelter_backend/internal/dto"
	"Pelter_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Route(app *fiber.App, gorm *gorm.DB) {
	repo := NewUserRepository(gorm)
	usecase := NewUserUsecase(repo)
	service := NewUserService(usecase)

	// group := app.Group("/users")
	// group.Get("/register", service.Register)
	app.Post("/register", middleware.ValidationMiddleware(&dto.RegisterRequest{}), service.Register)
}
