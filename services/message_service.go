package services

import "Marcketplace/model"

type MessageService interface {
	CreateConversation(convo model.Conversation) error
	SupprimerConversation(convoID int) error

	SendMessage(message model.Message) error

	GetMessageFromConversation(convoID int) ([]model.JMessage, error)
}
