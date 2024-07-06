package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Users struct {
	ID           int       `gorm:"primaryKey; index; AUTO_INCREMENT"`
	UserName     string    `gorm:"column:user_name; NOT NULL; unique; index"`
	Password     string    `gorm:"NOT NULL"`
	IsActive     bool      `gorm:"column:is_active"`
	CreateAt     time.Time `gorm:"column:create_at"`
	LastActiveAt time.Time `gorm:"column:last_active_at"`
	Image        string    `gorm:"column:path"`
	Email        string
}

type TokenClaims struct {
	UserName string `json:"user_name"`
	Image    string `json:"image"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}