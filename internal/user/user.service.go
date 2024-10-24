package user

import (
	"Pelter_backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

type (
	userService struct {
		userUsecase UserUsecase
	}
	UserService interface {
		InsertUser(ctx *fiber.Ctx) error
	}
)

func NewUserService(userUsecase UserUsecase) UserService {
	return &userService{
		userUsecase: userUsecase,
	}
}

func (r *userService) InsertUser(ctx *fiber.Ctx) error {
	_ = r.userUsecase.InsertUser(ctx.Context())

	return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
		Result: "successful",
	})
}
