package services

import (
	"Marcketplace/data/response"
	"Marcketplace/repository"

	"github.com/go-playground/validator/v10"
)

type ElementImp struct {
	ElementRepository repository.ElemRepository
	validate          *validator.Validate
}

func NewEleServiceImpl(ElementRepository repository.ElemRepository, validate *validator.Validate) Element {
	return &ElementImp{
		ElementRepository: ElementRepository,
		validate:          validate,
	}
}

// FindAllCategories implements Element.
func (e *ElementImp) FindAllCategories() ([]response.CategoryResponse, error) {
	result := e.ElementRepository.FindAllCategories()
	var truc []response.CategoryResponse
	if result == nil {
		return truc, nil
	}
	for _, value := range result {
		Category := response.CategoryResponse{
			ID:    value.ID,
			Title: value.Title,
			Img:   value.Image,
		}
		truc = append(truc, Category)
	}
	return truc, nil
}

// FindAllTags implements Element.
func (e *ElementImp) FindAllTags() ([]response.TagResponse, error) {
	result := e.ElementRepository.FindAllCategories()
	var Tags []response.TagResponse
	if result == nil {
		return Tags, nil
	}
	for _, value := range result {
		Category := response.TagResponse{
			ID:    value.ID,
			Title: value.Title,
		}
		Tags = append(Tags, Category)
	}
	return Tags, nil
}
