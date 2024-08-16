package services

import (
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"Marcketplace/model/objets"
)

type ObjService interface {
	Create(objet request.CreateObjRequest)
	Update(objet request.UpdateObjRequest)
	Delete(objetId int)
	FindById(objetId int) response.ObjResponse
	FindAll() []response.ObjResponse
	ObjByCategID(CID uint) ([]objets.Objects, error)
	ObjByArticleID(CID uint) (objets.Objects, error)
	GetArticles(CID uint, status string) ([]objets.Objects, error)
	FindByName(name string) ([]objets.Objects, error)
	FindByCategId(id int) ([]objets.Objects, error)
}
