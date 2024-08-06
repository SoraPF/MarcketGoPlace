package repository

import (
	"Marcketplace/helper"
	"Marcketplace/model"

	"gorm.io/gorm"
)

type MessageRepositoryImpl struct {
	Db *gorm.DB
}

func NewMessageRepositoryImpl(Db *gorm.DB) MessageRepository {
	return &MessageRepositoryImpl{Db: Db}
}

// CreateConversation implements MessageRepository.
func (m *MessageRepositoryImpl) CreateConversation(convo model.Conversation) {
	result := m.Db.Create(&convo)
	helper.ErrorPanic(result.Error)
}

// SupprimerConversation implements MessageRepository.
func (m *MessageRepositoryImpl) SupprimerConversation(convoID uint) {
	var convo model.Conversation
	result := m.Db.Where("id = ? ", convoID).Delete(&convo)
	helper.ErrorPanic(result.Error)
}

// SendMessage implements MessageRepository.
func (m *MessageRepositoryImpl) SendMessage(message model.Message) {
	result := m.Db.Create(&message)
	helper.ErrorPanic(result.Error)
}

// GetMessageFromConversation implements MessageRepository.
func (m *MessageRepositoryImpl) GetMessageFromConversation(convoID uint) []model.Message {
	var messages []model.Message
	result := m.Db.Where("id = ?", convoID).Find(&messages)
	helper.ErrorPanic(result.Error)
	return messages
}
