package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();unique" json:"uuid"`
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

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.UUID = uuid.New()
	return
}
