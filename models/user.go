package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	//ID       uint   `josn:"id" gorm : "primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"` // "admin" or "user"
}

// 🔐 Used only during login request
type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
