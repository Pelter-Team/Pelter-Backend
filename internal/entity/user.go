package entity

// Custom Enum Type
type roleType string

const (
	Customer   roleType = "customer"
	Admin      roleType = "admin"
	Foundation roleType = "foundation"
	Seller     roleType = "seller"
)

type User struct {
	ID             uint `gorm:"primary_key"`
	Name           string
	Surname        string
	Email          string `gorm:"unique"`
	Password       string
	PhoneNumber    *string  `gorm:"default:null"`
	ProfileURL     *string  `gorm:"default:null"`
	Role           roleType `gorm:"type:role_type;default:'customer'"`
	Address        *string  `gorm:"default:null"`
	Verified       bool     `gorm:"default:false"`
	FoundationName *string  `gorm:"default:null"`
	ProductID      []Product
	Transactions   []Transaction `gorm:"foreignKey:BuyerID"`
	Sales          []Transaction `gorm:"foreignKey:SellerID"`
	// `gorm:"foreignKey:UserID"` don't have to use as
	// gorm will automatically detect the foreign key
	// https://arc.net/l/quote/rwcyijew
}
