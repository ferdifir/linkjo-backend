package models

import "time"

type OrderDetail struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	OrderID   uint
	ProductID uint
	Quantity  int
	Subtotal  float64
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
