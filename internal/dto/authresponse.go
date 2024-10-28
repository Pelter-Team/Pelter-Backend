package dto

import (
	"github.com/golang-jwt/jwt/v5"
)

type RegisterRequest struct {
	Name           string  `json:"name" validate:"required"`
	Surname        string  `json:"surname" validate:"required"`
	Email          string  `json:"email" validate:"required,email"`
	Password       string  `json:"password" validate:"required,min=8"`
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
	jwt.RegisteredClaims
}

type AuthResponse struct { // response data for auth
	UserID      uint   `json:"user_id"`
	AccessToken string `json:"access_token"`
}
