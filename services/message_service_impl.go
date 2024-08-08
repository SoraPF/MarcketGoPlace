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
func (m *MsessageImp) CreateConversation(convo model.JConversation) (uint, error) {
	conv := model.Conversation{
		Name:   convo.Name,
		Seller: uint(convo.SellerID),
		Buyer:  uint(convo.BuyerID),
	}
	id, err := m.MessageRepository.CreateConversation(conv)
	if err != nil {
		return 0, err
	}
	return id, nil
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
		println("pas trouver")
		return 0, errors.New("pas trouver")
	}
	println("trouver")
	if checks.SellerID == checks.UserID && int(conversation.Seller) == checks.UserID {
		println("user is seller")
		return conversation.ID, nil
	}
	if int(conversation.Seller) != checks.SellerID && int(conversation.Buyer) != checks.UserID {
		println("user and buyer is ok")
		return 0, errors.New("le nom existe mais ce nest pas les bon utilisateur")
	}
	if conversation.ID == 0 {
		println("existe pas")
		return 0, errors.New("existe pas")
	}

	return conversation.ID, nil
}
