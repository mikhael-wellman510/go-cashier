package entities

import "time"

type User struct {
	Base
	Username    string `json:"username" gorm:"column:username;not null;unique"`
	Email       string `json:"email" gorm:"column:email;not null;unique"`
	Password    string `json:"password" gorm:"column:password;not null"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Address     string `json:"address" gorm:"column:address"`
	IsActive    bool   `json:"is_active" gorm:"column:is_active;default:false"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
type UserRequest struct {
	Username    string `json:"username" binding:"required" `
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type UserResponse struct {
	Id          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
