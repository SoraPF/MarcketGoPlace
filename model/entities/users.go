package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint
	Username string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	BirthDay string
	Address  string
	Phone    string
	NFAID    *uint
	NFA      *NFA `gorm:"foreignKey:NFAID;references:ID"`
	//Verified bool `gorm:"default:false"`
}
