package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Route(app *fiber.App, gorm *gorm.DB) {
	repo := NewUserRepository(gorm)
	usecase := NewUserUsecase(repo)
	service := NewUserService(usecase)

	group := app.Group("/users")
	group.Get("/", service.InsertUser)
}
