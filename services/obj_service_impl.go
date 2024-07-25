package services

import (
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"Marcketplace/helper"
	"Marcketplace/model/objets"
	"Marcketplace/repository"

	"github.com/go-playground/validator/v10"
)

type ObjServiceImpl struct {
	ObjRepository      repository.ObjectRepository
	StatusRepository   repository.StatusRepository
	CategoryRepository repository.CategoryRepository
	TagRepository      repository.TagRepository
	validate           *validator.Validate
}

func NewObjServiceImpl(objRepository repository.ObjectRepository, validate *validator.Validate) ObjService {
	return &ObjServiceImpl{
		ObjRepository: objRepository,
		validate:      validate,
	}
}

// Create implements ObjService.
func (o *ObjServiceImpl) Create(object request.CreateObjRequest) {
	err := o.validate.Struct(object)
	helper.ErrorPanic(err)
	var tags []objets.Tags
	for _, tagID := range object.Tags {
		tag := objets.Tags{ID: tagID}
		tags = append(tags, tag)
	}
	objModel := objets.Objects{
		IdVendeur:  object.IdVendeur,
		Title:      object.Title,
		Price:      object.Price,
		Desc:       object.Desc,
		StatusID:   object.StatusID,
		CategoryID: object.CategoryID,
		Tags:       tags,
	}
	o.ObjRepository.Save(objModel)
}

// Delete implements ObjService.
func (o *ObjServiceImpl) Delete(objetId int) {
	o.ObjRepository.Delete(objetId)
}

// FindAll implements ObjService.
func (o *ObjServiceImpl) FindAll() []response.ObjResponse {
	result := o.ObjRepository.FindAll()
	var objects []response.ObjResponse
	for _, value := range result {
		var tags []response.TagResponse
		for _, tag := range value.Tags {
			tags = append(tags, response.TagResponse{
				ID:    tag.ID,
				Title: tag.Title,
			})
		}
		object := response.ObjResponse{
			IdVendeur: value.IdVendeur,
			Title:     value.Title,
			Price:     value.Price,
			Desc:      value.Desc,
			Status: response.StatusResponse{
				ID:    value.Status.ID,
				Title: value.Status.Title,
			},
			Category: response.CategoryResponse{
				ID:    value.Category.ID,
				Title: value.Category.Title,
			},
			Tags: tags,
		}
		objects = append(objects, object)
	}
	return objects
}

// FindById implements ObjService.
func (o *ObjServiceImpl) FindById(objetId int) response.ObjResponse {
	objData, err := o.ObjRepository.FindById(objetId)
	helper.ErrorPanic(err)
	var tagsResponse []response.TagResponse
	var tags []response.TagResponse
	for _, tag := range objData.Tags {
		tags = append(tags, response.TagResponse{
			ID:    tag.ID,
			Title: tag.Title,
		})
	}
	ObjResponse := response.ObjResponse{
		IdVendeur: objData.IdVendeur,
		Title:     objData.Title,
		Price:     objData.Price,
		Desc:      objData.Desc,
		Status: response.StatusResponse{
			ID:    objData.Status.ID,
			Title: objData.Status.Title,
		},
		Category: response.CategoryResponse{
			ID:    objData.Category.ID,
			Title: objData.Category.Title,
		},
		Tags: tagsResponse,
	}
	return ObjResponse
}

// Update implements ObjService.
func (o *ObjServiceImpl) Update(objet request.UpdateObjRequest) {
	objData, err := o.ObjRepository.FindById(int(objet.ID))
	helper.ErrorPanic(err)

	objData.Title = objet.Title
	objData.Price = objet.Price
	objData.Desc = objet.Desc

	status, err := o.StatusRepository.FindById(objet.StatusID)
	helper.ErrorPanic(err)
	objData.Status = status

	category, err := o.CategoryRepository.FindById(objet.CategoryID)
	helper.ErrorPanic(err)
	objData.Category = category

	var tags []objets.Tags
	for _, tagID := range objet.Tags {
		tags = append(tags, objets.Tags{ID: tagID})
	}
	objData.Tags = tags
	o.ObjRepository.Update(objData)
}
func (o *ObjServiceImpl) ObjByCategID(CID uint) ([]objets.Objects, error) {
	LObject, err := o.ObjRepository.ObjByCategID(CID)
	if err != nil {
		return nil, err
	}

	return LObject, nil
}
