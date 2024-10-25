package user

import (
	"context"

	"gorm.io/gorm"
)

type (
	repository struct {
		Db *gorm.DB
	}

	UserRepository interface {
		RegisterUser(pctx context.Context) error
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{
		Db: db,
	}
}

func (r *repository) RegisterUser(pctx context.Context) error {
	return nil
}
