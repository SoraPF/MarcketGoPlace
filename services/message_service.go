package services

import "Marcketplace/model"

type MessageService interface {
	CreateConversation(convo model.JConversation) (uint, error)
	SupprimerConversation(convoID int) error

	SendMessage(message model.JMessage) error

	GetMessageFromConversation(convoID int) ([]model.JMessage, error)

	CheckMessenger(checks model.Checkids) (uint, error)
	GetListeMessageries(id int) ([]model.Conversation, error)
}
