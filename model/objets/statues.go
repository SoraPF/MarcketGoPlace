package objets

import "gorm.io/gorm"

type Statuses struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Title string `gorm:"not null"`
}
