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

	route := app.Group("/auth")
	route.Post("/register", middleware.ValidationMiddleware(&dto.RegisterRequest{}), service.Register)
	route.Post("/login", middleware.ValidationMiddleware(&dto.LoginRequest{}), service.Login)
	route.Get("/me", middleware.ValidateCookie, service.GetMe)
	route.Get("/logout", service.Logout)
	admin := app.Group("/admin")
	admin.Get("/users", middleware.ValidateCookie, service.GetUsers)
}
