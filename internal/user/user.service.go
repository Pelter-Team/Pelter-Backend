package user

import (
	"Pelter_backend/internal/dto"
	"context"

	"Pelter_backend/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type (
	userService struct {
		userUsecase UserUsecase
	}
	UserService interface {
		Register(ctx *fiber.Ctx) error
		// Login(ctx *fiber.Ctx) error
	}
)

func NewUserService(userUsecase UserUsecase) UserService {
	return &userService{
		userUsecase: userUsecase,
	}
}

func (s *userService) Register(ctx *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.HttpResponse{
			Error: "Invalid request format",
		})
	}

	user := entity.User{
		Name:           req.Name,
		Surname:        req.Surname,
		Email:          req.Email,
		Password:       req.Password, // pass plain pwd to use case
		PhoneNumber:    req.PhoneNumber,
		ProfileURL:     req.ProfileURL,
		Role:           entity.Customer,
		Address:        req.Address,
		Verified:       false,
		FoundationName: req.FoundationName,
	}

	if err := s.userUsecase.Register(context.Background(), &user); err != nil {
		if err.Error() == "email already registered" {
			return ctx.Status(fiber.StatusConflict).JSON(dto.HttpResponse{
				Error: "Email already registered",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error: "Failed to create user",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.HttpResponse{
		Result: "User registered successfully",
	})
}

// func (r *userService) Login(ctx *fiber.Ctx) error {

// }
