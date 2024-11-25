package entity

import "time"

type Transaction struct {
	ID        uint    `gorm:"primary_key"`
	ProductID uint    // linked to product table (ID)
	BuyerID   uint    // linked to user table (ID)
	SellerID  uint    // linked to user table (ID)
	Amount    float64 // This might be redundant -> getting from Product entity (Price)
	CreatedAt time.Time
	UpdatedAt time.Time
	Product   Product `gorm:"foreignKey:ProductID"`
	Buyer     User    `gorm:"foreignKey:BuyerID"`
	Seller    User    `gorm:"foreignKey:SellerID"`
}
