package services

import (
	"Marcketplace/model"
	"Marcketplace/repository"
	"errors"

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
func (m *MsessageImp) CreateConversation(convo model.JConversation) error {
	conv := model.Conversation{
		Name: convo.Name,
	}
	err := m.MessageRepository.CreateConversation(conv)
	if err != nil {
		return err
	}
	return nil
}

// supprimerConversation implements MessageService.
func (m *MsessageImp) SupprimerConversation(convoID int) error {
	err := m.MessageRepository.SupprimerConversation(uint(convoID))
	if err != nil {
		return err
	}
	return nil
}

// sendMessage implements MessageService.
func (m *MsessageImp) SendMessage(message model.JMessage) error {
	mess := model.Message{
		ConversationID: uint(message.ConversationID),
		SenderID:       uint(message.SenderID),
		Content:        message.Content,
	}
	err := m.MessageRepository.SendMessage(mess)
	if err != nil {
		return err
	}
	return nil
}

// GetMessageFromConversation implements MessageService.
func (m *MsessageImp) GetMessageFromConversation(convoID int) ([]model.JMessage, error) {
	messages, err := m.MessageRepository.GetMessageFromConversation(uint(convoID))
	if err != nil {
		return nil, err
	}
	var jmessages []model.JMessage
	for _, message := range messages {
		mess := model.JMessage{
			ConversationID: int(message.ConversationID),
			SenderID:       int(message.SenderID),
			Content:        message.Content,
		}
		jmessages = append(jmessages, mess)
	}

	return jmessages, nil
}

// CheckMessenger implements MessageService.
func (m *MsessageImp) CheckMessenger(checks model.Checkids) (uint, error) {
	conversation, err := m.MessageRepository.FindConversationByName(checks.Name)
	if err != nil {
		return 0, errors.New("pas trouver")
	}
	cpt := 0
	for _, user := range conversation.UserIDs {
		if int(user) == checks.SellerID || int(user) == checks.UserID {
			cpt++
		}
	}
	if cpt < 2 {
		return 0, errors.New("le nom existe mais ce nest pas les bon utilisateur")
	}

	return conversation.ID, nil
}
