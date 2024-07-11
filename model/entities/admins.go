package entities

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Id       uint
	Username string
	Email    string
	Password string
	NFAID    uint `gorm:"not null"`
	NFA      NFA  `gorm:"foreignKey:NFAID;references:ID"`
}
