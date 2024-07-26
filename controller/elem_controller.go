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

func (controller *ElemController) GetCategories() ([]response.CategoryResponse, error) {
	categs, err := controller.ElemService.FindAllCategories()
	if err != nil {
		return nil, err
	}
	return categs, nil
}

func (controller *ElemController) GetTags() ([]response.TagResponse, error) {
	tags, err := controller.ElemService.FindAllTags()
	if err != nil {
		return nil, err
	}
	return tags, nil
}
