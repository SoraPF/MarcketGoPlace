package controller

import (
	"Marcketplace/data/response"
	"Marcketplace/helper"
	"Marcketplace/model"
	"Marcketplace/services"
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
	err := mc.ms.CreateConversation(convo) //need a modification and shoul be created service
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

func (mc MessageController) GetMessageFromConversation(c *fiber.Ctx) []model.JMessage {
	id := c.Params("id")
	if id == "" {
		return nil
	}
	idInt, err := strconv.Atoi(id)
	helper.ErrorPanic(err)
	messages, err := mc.ms.GetMessageFromConversation(idInt)
	if err != nil {
		return nil
	}
	return messages
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
		return c.Status(fiber.StatusBadRequest).SendString("the body wasnt correct")
	}

	err := mc.ms.CheckMessenger(newMessage) //need a modification and shoul be created service
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Internal error the conversation couldnt be created")
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete notes data!",
	}
	return c.Status(fiber.StatusCreated).JSON(webResponse)

}
