package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint
	Username string
	Email    string `gorm:"uniqueIndex"`
	Password string
	NFAID    *uint
	NFA      NFA `gorm:"foreignKey:NFAID;references:ID"`
}
