package dto

type RegisterRequest struct {
	Name           string  `json:"name"`
	Surname        string  `json:"surname"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	PhoneNumber    *string `json:"phone_number,omitempty"`
	ProfileURL     string  `json:"profile_url"`
	Role           string  `json:"role"`
	Address        *string `json:"address,omitempty"`
	Verified       bool    `json:"verified"`
	FoundationName *string `json:"foundation_name,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
