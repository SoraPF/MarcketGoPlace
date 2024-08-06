package model

import (
	"Marcketplace/model/entities"
)

type Message struct {
	ID             uint          `gorm:"primaryKey;autoIncrement"`
	ConversationID uint          `gorm:"not null" json:"ConversationID"`
	Conversation   Conversation  `gorm:"foreignKey:ConversationID"`
	SenderID       uint          `gorm:"not null" json:"SenderID"`
	Sender         entities.User `gorm:"foreignKey:SenderID"`
	Content        string        `gorm:"not null" json:"Content"`
}

type Conversation struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null"`
}

type Participant struct {
	ID             uint          `gorm:"primaryKey;autoIncrement"`
	ConversationID uint          `gorm:"not null"`
	Conversation   Conversation  `gorm:"foreignKey:ConversationID"`
	UserID         uint          `gorm:"not null"`
	User           entities.User `gorm:"foreignKey:UserID"`
}
