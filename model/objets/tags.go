package objets

import "gorm.io/gorm"

type Tags struct {
	gorm.Model
	ID    uint
	Title string
}
