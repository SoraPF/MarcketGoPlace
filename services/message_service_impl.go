package services

import (
	"Marcketplace/model"
	"Marcketplace/repository"

	"github.com/go-playground/validator/v10"
)

type MsessageImp struct {
	MessageRepository repository.MessageRepository
	validate          *validator.Validate
}

func NewMesServiceImpl(MessageRepository repository.MessageRepository, validate *validator.Validate) MessageService {
	return &MsessageImp{
		MessageRepository: MessageRepository,
		validate:          validate,
	}
}

// createConversation implements MessageService.
func (m *MsessageImp) CreateConversation(convo model.Conversation) {
	if convo.Name != "" {
		m.MessageRepository.CreateConversation(convo)
	}
}

// supprimerConversation implements MessageService.
func (m *MsessageImp) SupprimerConversation(convoID int) {
	if convoID != 0 {
		m.MessageRepository.SupprimerConversation(uint(convoID))
	}
}

// sendMessage implements MessageService.
func (m *MsessageImp) SendMessage(message model.Message) {
	if message.ConversationID != 0 && message.SenderID != 0 && message.Conversation.ID != 0 {
		m.MessageRepository.SendMessage(message)
	}
}

// GetMessageFromConversation implements MessageService.
func (m *MsessageImp) GetMessageFromConversation(convoID int) []model.Message {
	var messages []model.Message
	if convoID != 0 {
		messages = m.MessageRepository.GetMessageFromConversation(uint(convoID))
	}
	return messages
}
