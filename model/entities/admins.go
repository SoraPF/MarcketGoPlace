package entities

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Id       uint
	Username string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	NFAID    *uint
	NFA      NFA `gorm:"foreignKey:NFAID;references:ID"`
}
