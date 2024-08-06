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

func (mc MessageController) GetMessageFromConversation(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("no such conversation")
	}
	idInt, err := strconv.Atoi(id)
	helper.ErrorPanic(err)
	messages, err := mc.ms.GetMessageFromConversation(idInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("didnt find conversation")
	}

	webResponse := map[string]interface{}{
		"code":     200,
		"status":   "ok",
		"message":  "you successful refuse the article!",
		"messages": messages,
	}
	return c.Status(fiber.StatusCreated).JSON(webResponse)
}
