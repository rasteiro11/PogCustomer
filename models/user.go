package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID       uint
	Email    string
	Password string
	Document string
}

type Credentials struct {
	Password string
	Email    string
}

type LoginRequest struct {
	*Credentials
}

type LoginResponse struct {
	Token     string
	ExpiresAt time.Time
}

type RegisterRequest struct {
	*Credentials
}

type RegisterResponse struct {
	Token     string
	ExpiresAt time.Time
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type ChangePasswordRequest struct {
	Password    string
	NewPassword string
	Email       string
}

type ChangePasswordResponse struct {
}

type RegisterUserRequest struct {
	Email    string
	Password string
	Document string
	Wallet   string
}

type RegisterUserResponse struct {
	Token     string
	ExpiresAt time.Time
}
