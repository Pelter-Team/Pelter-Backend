package entity

import "time"

type Product struct {
	ID             uint `gorm:"primary_key"`
	UserID         uint //linked to user table (ID)
	TransactionID  uint //linked to transection table (ID)
	ReviewID       uint //linked to review table (ID)
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
