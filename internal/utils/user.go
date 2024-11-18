package utils

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

func IsAdmin(ctx context.Context, db *gorm.DB, userId uint) (bool, error) {
	var user struct {
		Role string
	}
	if err := db.WithContext(ctx).Table("users").Select("role").Where("id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("user not found")
		}
		return false, err
	}

	return user.Role == "admin", nil
}

func IsOwner(ctx context.Context, db *gorm.DB, productId uint, userId uint) (bool, error) {
	var product struct {
		UserID uint
	}
	if err := db.WithContext(ctx).Table("products").Select("user_id").Where("id = ?", productId).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("product not found")
		}
		return false, err
	}

	return product.UserID == userId, nil
}
