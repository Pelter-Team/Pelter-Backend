package user

import (
	"Pelter_backend/internal/dto"
	"Pelter_backend/internal/entity"
	"Pelter_backend/internal/pkg/bcrypt"
	"Pelter_backend/internal/pkg/jwt"
	"context"
	"errors"

	"gorm.io/gorm"
)

type (
	userUsecase struct {
		userRepo UserRepository
	}

	UserUsecase interface {
		Register(pctx context.Context, user *entity.User) (dto.LoginResponse, string, error)
		Login(pctx context.Context, email, password string) (dto.LoginResponse, string, error)
	}
)

func NewUserUsecase(userRepo UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Register(pctx context.Context, user *entity.User) (dto.LoginResponse, string, error) {
	count, err := u.userRepo.CountUserByEmail(pctx, user.Email)
	if errors.Is(err, gorm.ErrDuplicatedKey) || count == 1 {
		return dto.LoginResponse{}, "", errors.New("email already registered")
	}

	if err != nil {
		return dto.LoginResponse{}, "", errors.New("failed to find user by email")
	}

	hashedPwd, _ := bcrypt.HashPassword(user.Password)
	user.Password = string(hashedPwd)

	if err := u.userRepo.Create(pctx, user); err != nil {
		return dto.LoginResponse{}, "", err
	}
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return dto.LoginResponse{}, "", err
	}

	// NOTE: return as need not the whole struct
	return dto.LoginResponse{
		UserID: user.ID,
		FirstName: user.Name,
		Surname: user.Surname,
		Email: user.Email,
		ProfileURL: *user.ProfileURL,
		Role: string(user.Role),
	}, token, nil
}

func (u *userUsecase) Login(pctx context.Context, email string, password string) (dto.LoginResponse, string, error) {
	user, err := u.userRepo.FindByEmail(pctx, email)
	if err != nil {
		return dto.LoginResponse{}, "", err
	}

	if !bcrypt.CheckPassword(user.Password, password) {
		return dto.LoginResponse{}, "", errors.New("invalid credentials")
	}
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return dto.LoginResponse{}, "", err
	}

	return dto.LoginResponse{
		UserID: user.ID,
		FirstName: user.Name,
		Surname: user.Surname,
		Email: user.Email,
		ProfileURL: *user.ProfileURL,
		Role: string(user.Role),
	}, token, nil
}
