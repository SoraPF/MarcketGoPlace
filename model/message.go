package model

import (
	"Marcketplace/model/entities"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID             uint          `gorm:"primaryKey;autoIncrement"`
	ConversationID uint          `gorm:"not null" json:"ConversationID"`
	Conversation   Conversation  `gorm:"foreignKey:ConversationID"`
	SenderID       uint          `gorm:"not null" json:"SenderID"`
	Sender         entities.User `gorm:"foreignKey:SenderID"`
	Content        string        `gorm:"not null" json:"Content"`
}
type JMessage struct {
	ConversationID int    `gorm:"not null" json:"ConversationID"`
	SenderID       int    `gorm:"not null" json:"SenderID"`
	Content        string `gorm:"not null" json:"Content"`
}

type Conversation struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null"`
}
type JConversation struct {
	ID   int
	Name string `json:"name"`
}

type Participant struct {
	gorm.Model
	ID             uint          `gorm:"primaryKey;autoIncrement"`
	ConversationID uint          `gorm:"not null"`
	Conversation   Conversation  `gorm:"foreignKey:ConversationID"`
	UserID         uint          `gorm:"not null"`
	User           entities.User `gorm:"foreignKey:UserID"`
}
