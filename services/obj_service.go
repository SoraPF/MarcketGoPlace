package services

import (
	"Marcketplace/data/request"
	"Marcketplace/data/response"
)

type ObjService interface {
	Create(objet request.CreateObjRequest)
	Update(objet request.UpdateObjRequest)
	Delete(objetId int)
	FindById(objetId int) response.ObjResponse
	FindAll() []response.ObjResponse
}
