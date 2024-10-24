package entity

import "time"

type Review struct {
	ID        uint `gorm:"primary_key"`
	ProductID uint 
	Rating    uint
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
