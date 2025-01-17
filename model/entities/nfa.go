package entities

import "gorm.io/gorm"

type NFA struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	QRcode string `gorm:"not null"`
	Secret string `gorm:"not null"`
}

func (u *User) IsNFA() bool {
	return u.NFAID != nil
}
