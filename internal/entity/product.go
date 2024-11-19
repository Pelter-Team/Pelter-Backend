package entity

import "time"

type Product struct {
	ID             uint `gorm:"primary_key"`
	UserID         uint
	ReviewID       []Review
	IsSold         bool
	Category       string
	Subcategory    string
	Description    string
	IsVerified     bool
	Price          float64
	ImageURL       string
	VaccineBookURL *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
