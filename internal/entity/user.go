package entity

import "time"

// Custom Enum Type
type roleType string

const (
	Customer   roleType = "customer"
	Admin      roleType = "admin"
	Foundation roleType = "foundation"
	Seller     roleType = "seller"
)

type User struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Surname     string
	Email       string
	Password    string
	PhoneNumber string
	ProfileURL  string
	// lineID         string
	Role           roleType `gorm:"role_type"`
	Address        string
	Verified       bool
	FoundationName *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ProductID      []Product
	Transactions   []Transaction `gorm:"foreignKey:BuyerID"`
	Sales          []Transaction `gorm:"foreignKey:SellerID"`
	// `gorm:"foreignKey:UserID"` don't have to use as
	// gorm will automatically detect the foreign key
	// https://arc.net/l/quote/rwcyijew
}
