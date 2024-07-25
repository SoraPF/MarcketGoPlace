package objets

import (
	"gorm.io/gorm"
)

type Objects struct {
	gorm.Model
	ID         uint       `gorm:"primaryKey;autoIncrement"`
	IdVendeur  int        `gorm:"not null"`
	Title      string     `gorm:"not null"`
	Price      int        `gorm:"not null"`
	Desc       string     `gorm:"not null"`
	StatusID   uint       `gorm:"not null"`
	Status     Statuses   `gorm:"foreignKey:StatusID"`
	CategoryID uint       `gorm:"not null"`
	Category   Categories `gorm:"foreignKey:CategoryID"`
	Tags       []Tags     `gorm:"many2many:object_tags"`
	Img        []string   `gorm:"type:json"`
}
