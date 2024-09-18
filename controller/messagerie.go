package controller

import (
	"Marcketplace/data/response"
	"Marcketplace/helper"
	"Marcketplace/model"
	"Marcketplace/services"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MessageController struct {
	ms services.MessageService
}

func NewMessController(service services.MessageService) *MessageController {
	return &MessageController{ms: service}
}

func (mc MessageController) CreateConversation(c *fiber.Ctx) error {
	var convo model.JConversation
	if err := c.BodyParser(&convo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("the body wasnt correct")
	}
	println("messenger body ok")
	id, err := mc.ms.CreateConversation(convo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("insternal error the conversation couldnt be created")
	}
	println("mensenger created")

	webResponse := map[string]interface{}{
		"code":    200,
		"status":  "ok",
		"message": "Login successful!",
		"id":      id,
	}
	return c.Status(fiber.StatusCreated).JSON(webResponse)
}
func (mc MessageController) SupprimerConversation(c *fiber.Ctx) error {
	type Conv struct {
		ID int `json:"id"`
	}
	var conv Conv
	if err := c.BodyParser(&conv); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("the body wasnt correct")
	}
	err := mc.ms.SupprimerConversation(conv.ID) //need a modification and shoul be created service
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("insternal error the conversation couldnt be created")
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete notes data!",
	}
	return c.Status(fiber.StatusCreated).JSON(webResponse)
}

func (mc MessageController) SendMessage(c *fiber.Ctx) error {
	var newMessage model.JMessage
	if err := c.BodyParser(&newMessage); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("the body wasnt correct")
	}

	err := mc.ms.SendMessage(newMessage) //need a modification and shoul be created service
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("insternal error the conversation couldnt be created")
	}

	/*ajouter la parti notification*/

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete notes data!",
	}
	return c.Status(fiber.StatusCreated).JSON(webResponse)

}

func (mc MessageController) GetMessageFromConversation(c *fiber.Ctx) ([]model.JMessage, error) {
	id := c.Params("id")
	if id == "" {
		return nil, nil
	}
	idInt, err := strconv.Atoi(id)
	if id == "" {
		return nil, err
	}
	messages, err := mc.ms.GetMessageFromConversation(idInt)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (mc MessageController) GetMessagesFromConversation(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("pas bon la requette")
	}
	idInt, err := strconv.Atoi(id)
	helper.ErrorPanic(err)
	messages, err := mc.ms.GetMessageFromConversation(idInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("pas bon la requette")
	}

	webResponse := map[string]interface{}{
		"code":     200,
		"status":   "ok",
		"message":  "Login successful!",
		"messages": messages, // Notez qu'on retourne le token sous forme de string
	}
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

func (mc MessageController) CheckMessenger(c *fiber.Ctx) error {

	var newMessage model.Checkids
	if err := c.BodyParser(&newMessage); err != nil {
		log.Printf("Error parsing body: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("the body wasnt correct")
	}

	conversationID, err := mc.ms.CheckMessenger(newMessage)
	if err != nil {
		if err.Error() != "pas trouver" {
			return c.Status(fiber.StatusInternalServerError).SendString("le nom existe mais ce nest pas les bon utilisateur")
		}
		webResponse := map[string]interface{}{
			"code":    200,
			"status":  "ok",
			"message": "Login successful!",
			"text":    "can be created",
		}
		return c.Status(fiber.StatusOK).JSON(webResponse)
	}

	webResponse := map[string]interface{}{
		"code":           200,
		"status":         "ok",
		"message":        "Login successful!",
		"conversationID": int(conversationID),
	}
	return c.Status(fiber.StatusCreated).JSON(webResponse)

}

func (u UserController) ProposePrice(c *fiber.Ctx) error {
	type proposePrice struct {
		Pprice   int    `json:"pPrice"`
		Oprice   int    `json:"oPrice"`
		Acheteur int    `json:"acheteur"`
		Vendeur  int    `json:"vendeur"`
		Aname    string `json:"Aname"`
	}

	var pp proposePrice
	if err := c.BodyParser(&pp); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	user := u.UserService.FindById(uint(pp.Acheteur))
	println("test", pp.Pprice, pp.Oprice, "v:", user.Name)
	pPriceStr := strconv.Itoa(pp.Pprice)
	vIdStr := strconv.Itoa(pp.Vendeur)
	oPriceStr := strconv.Itoa(pp.Oprice)
	aPriceStr := strconv.Itoa(pp.Acheteur)

	content := user.Name + " propose une offre de <strong>" + pPriceStr + "</strong> à la place de " +
		oPriceStr + " sur l'article " + pp.Aname + "."
	Notification("notification", vIdStr, aPriceStr, content, pp.Pprice)

	return c.SendStatus(fiber.StatusOK)
}

func (mc *MessageController) GetListeMessageries(id int) ([]model.Conversation, error) {
	var ListeMessageries []model.Conversation
	ListeMessageries, err := mc.ms.GetListeMessageries(id)
	if err != nil {
		return nil, err
	}
	for _, lm := range ListeMessageries {
		println(lm.Name)
	}
	return ListeMessageries, nil
}

func (mc *MessageController) GetConversation(id int) model.JConversation {
	cid := uint(id)
	conv := mc.ms.GetConversation(cid)
	return conv
}

func Notification(t string, uid string, nuid string, content string, price int) {
	if price != 0 {
		notification := Message{
			Type:        t,
			UserID:      uid,
			Uidnotifier: nuid,
			Content:     content,
			Price:       price,
		}
		broadcast <- notification
	} else {
		notification := Message{
			Type:        t,
			UserID:      uid,
			Uidnotifier: nuid,
			Content:     content,
		}
		broadcast <- notification
	}
}
