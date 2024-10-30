package user

import (
	"Pelter_backend/internal/entity"
	"Pelter_backend/internal/pkg/bcrypt"
	"Pelter_backend/internal/pkg/jwt"
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type (
	userUsecase struct {
		userRepo UserRepository
	}

	UserUsecase interface {
		Register(pctx context.Context, user *entity.User) error
		Login(pctx context.Context, email, password string) (*entity.User, string, error)
		Logout(ctx *fiber.Ctx) error
	}
)

func NewUserUsecase(userRepo UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Register(pctx context.Context, user *entity.User) error {

	existingUser, err := u.userRepo.FindByEmail(pctx, user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already registered")
	}

	hashedPwd, _ := bcrypt.HashPassword(user.Password)
	user.Password = string(hashedPwd)

	return u.userRepo.Create(pctx, user)
}

func (u *userUsecase) Login(pctx context.Context, email string, password string) (*entity.User, string, error) {
	user, err := u.userRepo.FindByEmail(pctx, email)
	if err != nil {
		return nil, "", err
	}

	if !bcrypt.CheckPassword(user.Password, password) {
		return nil, "", errors.New("invalid credentials")
	}
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (u *userUsecase) Logout(ctx *fiber.Ctx) error {
	_ = u.userRepo.Logout(ctx.Context())
	ctx.ClearCookie("access_token")
	return nil
}
