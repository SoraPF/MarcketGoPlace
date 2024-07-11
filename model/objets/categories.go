package objets

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	ID    uint
	Title string
}
