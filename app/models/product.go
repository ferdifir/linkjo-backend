package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID            uint    `gorm:"primaryKey;autoIncrement"`
	TenantID      uint    `gorm:"not null;index"`
	CategoryID    uint    `gorm:"index"`
	Name          string  `gorm:"type:varchar(255);not null"`
	SKU           *string `gorm:"type:varchar(50);unique"`
	Barcode       *string `gorm:"type:varchar(100);unique"`
	Price         float64 `gorm:"not null"`
	CostPrice     *float64
	Discount      float64        `gorm:"default:0"`
	Stock         int            `gorm:"not null;default:0"`
	MinStockAlert int            `gorm:"default:0"`
	Unit          string         `gorm:"type:varchar(50);not null;default:'pcs'"`
	SupplierID    *uint          `gorm:"index"`
	Image         string         `gorm:"type:text"`
	Description   string         `gorm:"type:text"`
	IsActive      bool           `gorm:"default:true"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (Product) TableName() string {
	return "products"
}
