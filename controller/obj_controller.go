package controller

import (
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"Marcketplace/helper"
	"Marcketplace/model/objets"
	"Marcketplace/services"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ObjController struct {
	objService services.ObjService
}

func NewObjController(service services.ObjService) *ObjController {
	return &ObjController{objService: service}
}

func (controller *ObjController) ObjCreate(ctx *fiber.Ctx) error {
	idVendeur := ctx.FormValue("id_vendeur")
	statusID := ctx.FormValue("status_id")
	title := ctx.FormValue("title")
	price := ctx.FormValue("price")
	desc := ctx.FormValue("desc")
	categoryID := ctx.FormValue("category_id")
	tags := ctx.FormValue("tags")

	idVendeurInt, err := strconv.Atoi(idVendeur)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id_vendeur"})
	}
	statusIDInt, err := strconv.Atoi(statusID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status_id"})
	}
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid price"})
	}
	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category_id"})
	}
	var tagInts []int
	if tags != "" {
		tagStrs := strings.Split(tags, ",")
		for _, tag := range tagStrs {
			tagInt, err := strconv.Atoi(tag)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag"})
			}
			tagInts = append(tagInts, tagInt)
		}
	}

	createObjRequest := request.CreateObjRequest{
		IdVendeur:  idVendeurInt,
		Title:      title,
		Price:      priceInt,
		Desc:       desc,
		StatusID:   statusIDInt,
		CategoryID: categoryIDInt,
		Tags:       tagInts,
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse form data"})
	}

	files := form.File["images"]
	if len(files) > 0 {
		for _, file := range files {
			filename := file.Filename
			filePath := filepath.Join("./public/img/product", filename)

			// Check if the directory exists, if not, create it
			if _, err := os.Stat("./public/img/product"); os.IsNotExist(err) {
				err := os.MkdirAll("./public/img/product", os.ModePerm)
				if err != nil {
					return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create directory"})
				}
			}

			// Try saving the file
			if err := ctx.SaveFile(file, filePath); err != nil {
				fmt.Printf("Failed to save file: %s\n", err)
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
			}
			createObjRequest.Img = append(createObjRequest.Img, filePath)
		}
	}

	controller.objService.Create(createObjRequest)

	webResponse := map[string]interface{}{
		"code":         200,
		"status":       "ok",
		"message":      "request successful wait until admin accepte!",
		"redirect_url": "/createOk",
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ObjController) ObjUpdate(ctx *fiber.Ctx) error {
	idVendeur := ctx.FormValue("id_vendeur")
	statusID := ctx.FormValue("status_id")
	title := ctx.FormValue("title")
	price := ctx.FormValue("price")
	desc := ctx.FormValue("desc")
	categoryID := ctx.FormValue("category_id")
	tags := ctx.FormValue("tags")

	// Add logging to check the received values
	log.Printf("Received values - id_vendeur: %s, status_id: %s, title: %s, price: %s, desc: %s, category_id: %s, tags: %s", idVendeur, statusID, title, price, desc, categoryID, tags)

	idVendeurInt, err := strconv.Atoi(idVendeur)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id_vendeur"})
	}
	statusIDInt, err := strconv.Atoi(statusID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status_id"})
	}
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid price"})
	}
	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category_id"})
	}
	var tagInts []int
	if tags != "" {
		tagStrs := strings.Split(tags, ",")
		for _, tag := range tagStrs {
			tagInt, err := strconv.Atoi(tag)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag"})
			}
			tagInts = append(tagInts, tagInt)
		}
	}

	updateObjRequest := request.UpdateObjRequest{
		IdVendeur:  idVendeurInt,
		Title:      title,
		Price:      priceInt,
		Desc:       desc,
		StatusID:   statusIDInt,
		CategoryID: categoryIDInt,
		Tags:       tagInts,
	}

	objId := ctx.Params("objId")
	id, err := strconv.Atoi(objId)
	helper.ErrorPanic(err)

	updateObjRequest.ID = id

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

func (controller *ObjController) GetArticles(CID uint, status string) ([]objets.Objects, error) {
	Object, err := controller.objService.GetArticles(CID, status)
	if err != nil {
		return []objets.Objects{}, err
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
