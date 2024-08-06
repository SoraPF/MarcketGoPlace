package controller

import (
	"Marcketplace/data/response"
	"Marcketplace/model"

	"github.com/gofiber/fiber/v2"
)

func CreateConversation(c *fiber.Ctx) error {
	var convo model.Conversation
	if err := c.BodyParser(&convo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("the body wasnt correct")
	}
	conv, err := controller.service.createConversation() //need a modification and shoul be created service
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
func supprimerConversation(c *fiber.Ctx) error {
	type Conv struct {
		ID int `json:"id"`
	}
	var conv Conv
	if err := c.BodyParser(&conv); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("the body wasnt correct")
	}
	conv, err := controller.service.supprimerConversation() //need a modification and shoul be created service
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

func sendMessage(c *fiber.Ctx) error {
	var newMessage model.Message
	if err := c.BodyParser(&newMessage); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("the body wasnt correct")
	}
	participants, err := controller.service.findConversation() //need a modification and shoul be created service
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("insternal error the conversation couldnt be created")
	}

	message, err := controller.service.sendMessage() //need a modification and shoul be created service
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("insternal error the conversation couldnt be created")
	}

	for participant := range participants {
		if participant.ID != newMessage.SenderID {
			messageNotification(c, participant.ID)
		}
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete notes data!",
	}
	return c.Status(fiber.StatusCreated).JSON(webResponse)

}

/*
func messenger(ctx *fiber.Ctx) {

}
*/
