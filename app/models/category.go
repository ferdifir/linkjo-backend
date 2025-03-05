package models

type Categories struct {
	ID       uint         `gorm:"primaryKey;autoIncrement"`
	Name     string       `gorm:"type:varchar(255);not null"`
	ParentID *uint        `gorm:"index"`
	Children []Categories `gorm:"foreignKey:ParentID"`
}
