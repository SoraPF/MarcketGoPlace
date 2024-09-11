package repository

import (
	"Marcketplace/helper"
	"Marcketplace/model"
	"errors"

	"gorm.io/gorm"
)

type MessageRepositoryImpl struct {
	Db *gorm.DB
}

func NewMessageRepositoryImpl(Db *gorm.DB) MessageRepository {
	return &MessageRepositoryImpl{Db: Db}
}

// CreateConversation implements MessageRepository.
func (m *MessageRepositoryImpl) CreateConversation(convo model.Conversation) (uint, error) {
	result := m.Db.Create(&convo)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
		return 0, result.Error
	}
	return convo.ID, nil
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

func (m *MessageRepositoryImpl) FindConversationByName(name string) (*model.Conversation, error) {
	var convo model.Conversation
	result := m.Db.Where("name LIKE ?", name).Find(&convo)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &convo, nil
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
	result := m.Db.Where("conversation_id = ?", convoID).Find(&messages)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
		return messages, result.Error
	}
	return messages, nil
}

func (m *MessageRepositoryImpl) GetListeMessageries(convoID uint) ([]model.Conversation, error) {
	var messages []model.Conversation
	result := m.Db.Where("seller = ? OR buyer = ?", convoID, convoID).Find(&messages)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
		return messages, result.Error
	}
	return messages, nil
}
