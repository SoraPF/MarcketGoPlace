package entities

import "gorm.io/gorm"

type NFA struct {
	gorm.Model
	ID     uint
	QRcode string
	Secret string
}
