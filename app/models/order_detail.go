package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderDetail struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	OrderID   uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	Quantity  int  `gorm:"not null"`
	Subtotal  float64
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Order   Order   `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
	Product Product `gorm:"foreignKey:ProductID"`
}
