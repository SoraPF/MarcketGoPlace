package repository

import "Marcketplace/model"

type MessageRepository interface {
	CreateConversation(convo model.Conversation)
	SupprimerConversation(convoID uint)

	SendMessage(message model.Message)

	GetMessageFromConversation(convoID uint) []model.Message
}
