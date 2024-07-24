package objets

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Title string `gorm:"not null"`
	Image string
}
