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

type Conversation struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Name   string `gorm:"not null"`
	Seller uint   `gorm:"not null"`
	Buyer  uint   `gorm:"not null"`
}

type Participant struct {
	gorm.Model
	ID             uint          `gorm:"primaryKey;autoIncrement"`
	ConversationID uint          `gorm:"not null"`
	Conversation   Conversation  `gorm:"foreignKey:ConversationID"`
	UserID         uint          `gorm:"not null"`
	User           entities.User `gorm:"foreignKey:UserID"`
}

type JMessage struct {
	ConversationID int    `gorm:"not null" json:"ConversationID"`
	SenderID       int    `gorm:"not null" json:"SenderID"`
	Content        string `gorm:"not null" json:"Content"`
}

type JConversation struct {
	ID       int
	Name     string `json:"Name"`
	SellerID int    `json:"SellerID"`
	BuyerID  int    `json:"BuyerID"`
}

type Checkids struct {
	UserID   int    `json:"UserID"`
	SellerID int    `json:"SellerID"`
	Name     string `json:"Name"`
}
