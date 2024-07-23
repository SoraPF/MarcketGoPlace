package objets

import "gorm.io/gorm"

type Tags struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Title string `gorm:"not null"`
}
