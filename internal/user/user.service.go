package user

import (
	"Pelter_backend/internal/dto"
	"errors"

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
		GetUsers(ctx *fiber.Ctx) error
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
			Error:   err.Error(),
			Success: false,
		})
	}
	role := entity.RoleType(req.Role)
	user := entity.User{
		Name:           req.Name,
		Surname:        req.Surname,
		Email:          req.Email,
		Password:       req.Password, // pass plain pwd to use case
		PhoneNumber:    req.PhoneNumber,
		ProfileURL:     req.ProfileURL,
		Role:           role,
		Address:        req.Address,
		Verified:       false,
		FoundationName: req.FoundationName,
	}

	userRes, token, err := s.userUsecase.Register(ctx.UserContext(), &user)

	if err != nil {
		if err.Error() == "email already registered" { // check error from use case
			return ctx.Status(fiber.StatusConflict).JSON(dto.HttpResponse{
				Error:   "Email already registered",
				Success: false,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   "Failed to create user",
			Success: false,
		})
	}

	utils.SetCookie(ctx, "access_token", token)

	return ctx.Status(fiber.StatusCreated).JSON(dto.HttpResponse{
		Result:  userRes,
		Success: true,
	})
}

var (
	ErrInvalidCredential = errors.New("invalid credentials")
)

func (s *userService) Login(ctx *fiber.Ctx) error {

	req := ctx.Locals("body").(*dto.LoginRequest)

	user, token, err := s.userUsecase.Login(ctx.UserContext(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, ErrInvalidCredential) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
				Error:   "Invalid email or password",
				Success: false,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Error:   "Failed to login" + err.Error(),
			Success: false,
		})
	}

	utils.SetCookie(ctx, "access_token", token) // set token in cookie

	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  user,
		Success: true,
	})
}

func (s *userService) Logout(ctx *fiber.Ctx) error {
	ctx.ClearCookie("access_token")
	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  "Logged out successfully",
		Success: true,
	})
}

func (s *userService) GetUsers(ctx *fiber.Ctx) error {
	users, err := s.userUsecase.GetUsers(ctx.UserContext())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
			Result:  dto.UserResponse{},
			Error:   err.Error(),
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result:  users,
		Success: true,
	})
}
