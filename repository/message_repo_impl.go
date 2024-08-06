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
func (m *MessageRepositoryImpl) CreateConversation(convo model.Conversation) error {
	result := m.Db.Create(&convo)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
		return result.Error
	}
	return nil
}

// SupprimerConversation implements MessageRepository.
func (m *MessageRepositoryImpl) SupprimerConversation(convoID uint) error {
	var convo model.Conversation
	result := m.Db.Where("id = ? ", convoID).Delete(&convo)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
		return result.Error
	}
	return nil
}

// SendMessage implements MessageRepository.
func (m *MessageRepositoryImpl) SendMessage(message model.Message) error {
	result := m.Db.Create(&message)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
		return result.Error
	}
	return nil
}

// GetMessageFromConversation implements MessageRepository.
func (m *MessageRepositoryImpl) GetMessageFromConversation(convoID uint) ([]model.Message, error) {
	var messages []model.Message
	result := m.Db.Where("id = ?", convoID).Find(&messages)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
		return messages, result.Error
	}
	return messages, nil
}
