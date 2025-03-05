package models

import "time"

type Order struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	UserID        uint   `gorm:"not null"`
	CustomerName  string `gorm:"type:varchar(255);not null"`
	TableNumber   string `gorm:"type:varchar(50);not null"`
	TotalPrice    float64
	PaymentStatus string    `gorm:"type:varchar(50);not null"`
	PaymentMethod string    `gorm:"type:varchar(50);not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
