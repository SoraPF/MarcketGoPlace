package controller

import (
	"Marcketplace/data/response"
	"Marcketplace/services"
)

type ElemController struct {
	ElemService services.Element
}

func NewElemController(service services.Element) *ElemController {
	return &ElemController{ElemService: service}
}

func (controller *ElemController) GetCategories() response.AllCategory {
	categs, err := controller.ElemService.FindAllCategories()
	if err != nil {
		return response.AllCategory{}
	}
	return categs
}

func (controller *ElemController) GetTags() response.AllTags {
	categs, err := controller.ElemService.FindAllTags()
	if err != nil {
		return response.AllTags{}
	}
	return categs
}
