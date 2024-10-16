package entity

import "time"

type Transaction struct {
	ID        uint `gorm:"primary_key"`
	BuyerID   uint //linked to user table (ID)
	SellerID  uint //linked to user table (ID)
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
