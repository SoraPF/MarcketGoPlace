package controller

import (
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"Marcketplace/helper"
	"Marcketplace/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NoteController struct {
	noteService services.NoteService
}

func NewNoteController(service services.NoteService) *NoteController {
	return &NoteController{noteService: service}
}

func (controller *NoteController) Create(ctx *fiber.Ctx) error {
	updateNoteRequest := request.CreateNoteRequest{}
	err := ctx.BodyParser(&updateNoteRequest)
	helper.ErrorPanic(err)
	controller.noteService.Create(updateNoteRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully created notes data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *NoteController) Update(ctx *fiber.Ctx) error {
	updateNoteRequest := request.UpdateNoteRequest{}
	err := ctx.BodyParser(&updateNoteRequest)
	helper.ErrorPanic(err)

	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	updateNoteRequest.Id = id

	controller.noteService.Update(updateNoteRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully update notes data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *NoteController) Delete(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)
	controller.noteService.Delete(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete notes data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
func (controller *NoteController) FindById(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	NoteController := controller.noteService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete notes data!",
		Data:    NoteController,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)

}
func (controller *NoteController) FindAll(ctx *fiber.Ctx) error {
	NoteController := controller.noteService.FindAll()

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete notes data!",
		Data:    NoteController,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
