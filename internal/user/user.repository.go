package user

import (
	"context"

	"gorm.io/gorm"

	"Pelter_backend/internal/entity"
)

type (
	userRepository struct {
		Db *gorm.DB
	}

	UserRepository interface {
		Create(pctx context.Context, user *entity.User) error
		FindByEmail(pctx context.Context, email string) (*entity.User, error)
		Logout(pctx context.Context) error
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		Db: db,
	}
}

func (r *userRepository) Create(pctx context.Context, user *entity.User) error {
	return r.Db.WithContext(pctx).Create(user).Error
}

func (r *userRepository) FindByEmail(pctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := r.Db.WithContext(pctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(pctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	if err := r.Db.WithContext(pctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Logout(pctx context.Context) error {
	return nil
}
