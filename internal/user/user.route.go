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
	app.Post("/register", middleware.ValidationMiddleware(&dto.RegisterRequest{}), service.Register)
	app.Post("/login", middleware.ValidationMiddleware(&dto.LoginRequest{}), service.Login)
	app.Get("/logout", middleware.ValidateCookie, service.Logout)
}
