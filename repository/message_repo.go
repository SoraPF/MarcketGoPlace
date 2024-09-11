package repository

import "Marcketplace/model"

type MessageRepository interface {
	CreateConversation(convo model.Conversation) (uint, error)
	SupprimerConversation(convoID uint) error

	SendMessage(message model.Message) error
	GetMessageFromConversation(convoID uint) ([]model.Message, error)
	FindConversationByName(name string) (*model.Conversation, error)
	GetListeMessageries(convoID uint) ([]model.Conversation, error)
}
