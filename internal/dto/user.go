package dto

import "time"

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
}

type LoginResponse struct {
	UserID      uint   `json:"user_id"`
	AccessToken string `json:"access_token"`
}

type CookieRequest struct {
	Name        string    `cookie:"name"`
	Value       string    `cookie:"value"`
	Path        string    `cookie:"path"`
	Domain      string    `cookie:"domain"`
	MaxAge      int       `cookie:"max_age"`
	Expires     time.Time `cookie:"expires"`
	Secure      bool      `cookie:"secure"`
	HTTPOnly    bool      `cookie:"http_only"`
	SameSite    string    `cookie:"same_site"`
	SessionOnly bool      `cookie:"session_only"`
}