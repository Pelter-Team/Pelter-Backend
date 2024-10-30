package user

import (
	"Pelter_backend/internal/dto"

	"context"

	"Pelter_backend/internal/entity"

	"github.com/gofiber/fiber/v2"

	"Pelter_backend/internal/utils"
)

type (
	userService struct {
		userUsecase UserUsecase
	}
	UserService interface {
		Register(ctx *fiber.Ctx) error
		Login(ctx *fiber.Ctx) error
		Logout(ctx *fiber.Ctx) error
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
			Error: "fuck " + err.Error(),
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
		if err.Error() == "email already registered" { // check error from use case
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

func (s *userService) Login(ctx *fiber.Ctx) error {

	req := ctx.Locals("body").(*dto.LoginRequest)

	user, token, err := s.userUsecase.Login(context.Background(), req.Email, req.Password)
	if err != nil {
		if err.Error() == "invalid credentials" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
				Error: "Invalid email or password",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error: "Failed to login" + err.Error(),
		})
	}

	utils.SetCookie(ctx, "access_token", token) // set token in cookie

	return ctx.Status(fiber.StatusOK).JSON(dto.LoginResponse{
		UserID:      user.ID,
		AccessToken: "Created",
	})
}

func (s *userService) Logout(ctx *fiber.Ctx) error {
	err := s.userUsecase.Logout(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error: "Failed to logout",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result: "Logged out successfully",
	})
}
