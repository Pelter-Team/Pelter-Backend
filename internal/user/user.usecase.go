package user

import (
	"context"
)

type (
	usecase struct {
		userRepo UserRepository
	}

	UserUsecase interface {
		InsertUser(pctx context.Context) error
	}
)

func NewUserUsecase(userRepo UserRepository) UserUsecase {
	return &usecase{
		userRepo: userRepo,
	}
}

func (r *usecase) InsertUser(pctx context.Context) error {
	_ = r.userRepo.RegisterUser(pctx)
	return nil
}
