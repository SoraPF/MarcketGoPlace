package repository

import "Marcketplace/model"

type MessageRepository interface {
	CreateConversation(convo model.Conversation) error
	SupprimerConversation(convoID uint) error

	SendMessage(message model.Message) error

	GetMessageFromConversation(convoID uint) ([]model.Message, error)
}
