package entity

import "time"

type Review struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint //linked to user table (ID)
	ProductID uint //linked to product table (ID)
	Rating    uint
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
