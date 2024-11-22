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
		FindByEmail(pctx context.Context, email string) (entity.User, error)
		CountUserByEmail(pctx context.Context, email string) (int64, error)
		FindByID(pctx context.Context, id uint) (*entity.User, error)
		GetUsers(pctx context.Context) ([]*entity.User, error)
	}
)

func (r *userRepository) userTable(pctx context.Context) *gorm.DB {
	return r.Db.Table("users").WithContext(pctx)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		Db: db,
	}
}

func (r *userRepository) Create(pctx context.Context, user *entity.User) error {
	return r.userTable(pctx).Create(user).Error
}

func (r *userRepository) FindByEmail(pctx context.Context, email string) (entity.User, error) {
	var user entity.User
	if err := r.userTable(pctx).Where("email = ?", email).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) CountUserByEmail(pctx context.Context, email string) (int64, error) {
	count := new(int64)
	if err := r.userTable(pctx).Where("email = ?", email).Count(count).Error; err != nil {
		return -1, err
	}
	return *count, nil
}

func (r *userRepository) FindByID(pctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	if err := r.userTable(pctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetUsers(pctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	if err := r.userTable(pctx).Where("role != ?", entity.Admin).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
