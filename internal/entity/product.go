package entity

import (
	"Pelter_backend/internal/dto"
	"time"
)

type Product struct {
	ID             uint `gorm:"primary_key"`
	UserID         uint
	User           User `gorm:"foreignKey:UserID"`
	Review         []Review
	Name           string
	IsSold         bool
	Category       string
	Subcategory    string
	Description    string
	IsVerified     bool
	Price          float64
	ImageURL       string
	VaccineBookURL *string
	Phone          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (pr *Product) ConvertToProductResponse() dto.ProductResponse {
	return dto.ProductResponse{
		ID:             pr.ID,
		Name:           pr.Name,
		Owner:          pr.User.Name,
		Role:           pr.User.Role.String(),
		UserID:         pr.UserID,
		Phone:          pr.User.PhoneNumber,
		UserProfileUrl: pr.User.ProfileURL,
		IsSold:         pr.IsSold,
		Category:       pr.Category,
		Subcategory:    pr.Subcategory,
		Description:    pr.Description,
		IsVerified:     pr.IsVerified,
		Price:          pr.Price,
		ImageURL:       pr.ImageURL,
		VaccineBookURL: pr.VaccineBookURL,
		CreatedAt:      pr.CreatedAt,
		UpdatedAt:      pr.UpdatedAt,
	}
}
