package objets

import "gorm.io/gorm"

type Statuses struct {
	gorm.Model
	ID    uint
	Title string
}
