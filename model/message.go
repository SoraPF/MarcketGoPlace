package model

import "Marcketplace/model/entities"

type Messages struct {
	ID              uint          `gorm:"primaryKey;autoIncrement"`
	Conversation_id uint          `gorm:"not null"`
	Conversation    Conversation  `gorm:"foreignKey:Userid"`
	Sender_id       uint          `gorm:"not null"`
	Sender          entities.User `gorm:"foreignKey:Userid"`
	Content         string        `gorm:"not null"`
}

type Conversation struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null"`
}

type Participant struct {
	ID              uint          `gorm:"primaryKey;autoIncrement"`
	Conversation_id uint          `gorm:"not null"`
	Conversation    Conversation  `gorm:"foreignKey:Userid"`
	User_id         uint          `gorm:"not null"`
	User            entities.User `gorm:"foreignKey:Userid"`
}
