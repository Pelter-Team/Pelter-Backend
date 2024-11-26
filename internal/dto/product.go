package dto

import "time"

type ProductRequest struct {
	Name           string  `json:"name" validate:"required"`
	IsSold         bool    `json:"is_sold" validate:"required"`
	Category       string  `json:"category" validate:"required"`
	Subcategory    string  `json:"subcategory" validate:"required"`
	Description    string  `json:"description" validate:"required"`
	IsVerified     bool    `json:"is_verified" validate:"required"`
	Price          float64 `json:"price" validate:"required"`
	ImageURL       string  `json:"image_url" validate:"required,url"`
	VaccineBookURL *string `json:"vaccine_book_url,omitempty" validate:"omitempty,url"`
}

type ProductUpdateRequest struct {
	Name           string  `json:"name"`
	IsSold         bool    `json:"is_sold"`
	Category       string  `json:"category"`
	Subcategory    string  `json:"subcategory"`
	Description    string  `json:"description"`
	IsVerified     bool    `json:"is_verified"`
	Price          float64 `json:"price"`
	ImageURL       string  `json:"image_url"`
	VaccineBookURL *string `json:"vaccine_book_url,omitempty"`
}

type ProductUpdateVerificationStatus struct {
	IsVerified bool `json:"is_verified"`
}

type ProductResponse struct {
	ID             uint      `json:"id"`
	UserID         uint      `json:"user_id"`
	Owner          string    `json:"owner"`
	UserProfileUrl *string   `json:"user_profile_url"`
	Role           string    `json:"role"`
	TransactionID  uint      `json:"transaction_id"`
	Phone          *string   `json:"phone"`
	Review         []uint    `json:"review_id"`
	Name           string    `json:"name" validate:"required"`
	IsSold         bool      `json:"is_sold" validate:"required"`
	Category       string    `json:"category" validate:"required"`
	Subcategory    string    `json:"subcategory" validate:"required"`
	Description    string    `json:"description" validate:"required"`
	IsVerified     bool      `json:"is_verified" validate:"required"`
	Price          float64   `json:"price" validate:"required"`
	ImageURL       string    `json:"image_url" validate:"required,url"`
	VaccineBookURL *string   `json:"vaccine_book_url" validate:"omitempty,url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ProductWithUserResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Name        string    `json:"name" validate:"required"`
	IsSold      bool      `json:"is_sold" validate:"required"`
	IsVerified  bool      `json:"is_verified" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description" validate:"required"`
	Category    string    `json:"category" validate:"required"`
	Subcategory string    `json:"subcategory" validate:"required"`
}

type UpdateProductIsSold struct {
	IsSold bool `json:"is_sold"`
}

type QueryProductsIn struct {
	ProductsId []uint `json:"productIds" validate:"required"`
}
