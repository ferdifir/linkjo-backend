package models

import (
	"time"
)

type User struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	OwnerName  string    `json:"owner_name"`
	OutletName string    `json:"outlet_name"`
	Email      string    `json:"email" gorm:"unique"`
	Password   string    `json:"-"`
	Phone      string    `json:"phone" gorm:"unique"`
	Role       string    `json:"role" gorm:"default:user"`
	Photo      string    `json:"photo" gorm:"default:'default.png'"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	IsActive   bool      `json:"is_active" gorm:"default:true"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
