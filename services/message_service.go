package services

import "Marcketplace/model"

type MessageService interface {
	CreateConversation(convo model.Conversation)
	SupprimerConversation(convoID int)

	SendMessage(message model.Message)

	GetMessageFromConversation(convoID int) []model.Message
}
