package entities

import "gorm.io/gorm"

type NFA struct {
	gorm.Model
	ID     uint
	QRcode string
	Secret string
}

func (u *User) IsNFA() bool {
	return u.NFA != nil && u.NFA.Secret != ""
}
