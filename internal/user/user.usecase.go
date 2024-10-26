package user

import (
	"Pelter_backend/internal/entity"
	"context"
	"errors"

	"Pelter_backend/internal/pkg/bcrypt"
)

type (
	userUsecase struct {
		userRepo UserRepository
	}

	UserUsecase interface {
		Register(pctx context.Context, user *entity.User) error
	}
)

func NewUserUsecase(userRepo UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Register(ctx context.Context, user *entity.User) error {

	existingUser, err := u.userRepo.FindByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already registered")
	}

	hashedPwd, _ := bcrypt.HashPassword(user.Password)
	user.Password = string(hashedPwd)

	return u.userRepo.Create(ctx, user)
}
