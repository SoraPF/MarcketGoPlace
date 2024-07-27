package controller

import (
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"Marcketplace/helper"
	"Marcketplace/model/objets"
	"Marcketplace/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ObjController struct {
	objService services.ObjService
}

func NewObjController(service services.ObjService) *ObjController {
	return &ObjController{objService: service}
}

func (controller *ObjController) ObjCreate(ctx *fiber.Ctx) error {
	createObjRequest := request.CreateObjRequest{}
	err := ctx.BodyParser(&createObjRequest)
	helper.ErrorPanic(err)
	controller.objService.Create(createObjRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully created notes data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ObjController) ObjUpdate(ctx *fiber.Ctx) error {
	updateObjRequest := request.UpdateObjRequest{}
	err := ctx.BodyParser(&updateObjRequest)
	helper.ErrorPanic(err)

	objId := ctx.Params("objId")
	id, err := strconv.ParseUint(objId, 10, 32)
	helper.ErrorPanic(err)

	updateObjRequest.ID = uint(id)

	controller.objService.Update(updateObjRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully update objs data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ObjController) ObjDelete(ctx *fiber.Ctx) error {
	objId := ctx.Params("objId")
	id, err := strconv.Atoi(objId)
	helper.ErrorPanic(err)
	controller.objService.Delete(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete objs data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ObjController) ObjFindById(ctx *fiber.Ctx) error {
	objId := ctx.Params("objId")
	id, err := strconv.Atoi(objId)
	helper.ErrorPanic(err)

	ObjController := controller.objService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete objects data!",
		Data:    ObjController,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ObjController) ObjFindAll(ctx *fiber.Ctx) error {
	ObjController := controller.objService.FindAll()

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete objects data!",
		Data:    ObjController,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ObjController) ObjByCategID(CID uint) []objets.Objects {
	LObject, err := controller.objService.ObjByCategID(CID)
	if err != nil {
		return nil
	}
	return LObject
}

func (controller *ObjController) ObjByArticleID(CID uint) (objets.Objects, error) {
	Object, err := controller.objService.ObjByArticleID(CID)
	if err != nil {
		return objets.Objects{}, err
	}
	return Object, nil
}

func RequestCreateArticle(ctx *fiber.Ctx) error {
	req := request.CreateObjRequest{}
	err := ctx.BodyParser(&req)
	helper.ErrorPanic(err)

	//NotifiedAdminNewArticle(ctx, &req)

	webResponse := map[string]interface{}{
		"code":         200,
		"status":       "ok",
		"message":      "request successful wait until admin accepte!",
		"redirect_url": "/createOk",
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ObjController) AdminResponceNewArticle(ctx *fiber.Ctx) error {
	req := request.CreateObjRequest{}
	err := ctx.BodyParser(&req)
	helper.ErrorPanic(err)
	if req.Title == "" || req.Price == 0 || req.Desc == "" || req.CategoryID == 0 || req.IdVendeur == 0 || req.StatusID == 0 || req.Tags == nil {
		webResponse := map[string]interface{}{
			"code":    200,
			"status":  "ok",
			"message": "you successful refuse the article!",
		}
		return ctx.Status(fiber.StatusOK).JSON(webResponse)
	}

	//NotifiedUserNewArticle(ctx, &req)
	//return controller.ObjCreate(ctx)

	webResponse := map[string]interface{}{
		"code":    200,
		"status":  "ok",
		"message": "you successful accepted the article!",
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
