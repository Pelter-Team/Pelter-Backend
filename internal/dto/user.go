package dto

import (
	"time"
)

type RegisterRequest struct {
	Name           string  `json:"name" validate:"required"`
	Surname        string  `json:"surname" validate:"required"`
	Email          string  `json:"email" validate:"required,email"`
	Password       string  `json:"password" validate:"required,min=4"`
	PhoneNumber    *string `json:"phone_number,omitempty" validate:"omitempty"`
	ProfileURL     *string `json:"profile_url,omitempty" validate:"omitempty,url"`
	Role           string  `json:"role" validate:"required,oneof=admin customer foundation seller"`
	Address        *string `json:"address,omitempty" validate:"omitempty"`
	Verified       bool    `json:"verified" validate:"omitempty"`
	FoundationName *string `json:"foundation_name,omitempty" validate:"omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	UserID     uint    `json:"userId"`
	FirstName  string  `json:"firstname"`
	Surname    string  `json:"surname"`
	Email      string  `json:"email"`
	ProfileURL *string `json:"profileUrl"`
	Role       string  `json:"role"`
	Phone      *string `json:"phone"`
}

type UserResponse struct {
	CreatedAt time.Time `json:"created_at"`
	Address   *string   `json:"address"`
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	Phone     *string   `json:"phone"`
	Role      string    `json:"role"`
}

type MeResponse struct {
	UserID     uint    `json:"user_id"`
	Username   string  `json:"username"`
	Role       string  `json:"role"`
	ProfileURL *string `json:"profileUrl"`
	Surname    string  `json:"surname"`
	Phone      *string `json:"phone"`
}
