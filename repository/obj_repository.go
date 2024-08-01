package repository

import "Marcketplace/model/objets"

type ObjectRepository interface {
	Save(object objets.Objects)
	Update(object objets.Objects)
	Delete(objectId int)
	FindById(objectId int) (objets.Objects, error)
	FindAll() []objets.Objects
	ObjByCategID(CID uint) ([]objets.Objects, error)
	ObjByArticleID(CID uint) (objets.Objects, error)
	GetArticles(CID uint, status string) (objets.Objects, error)
}

type StatusRepository interface {
	FindById(StatusId uint) (objets.Statuses, error)
}

type CategoryRepository interface {
	FindById(CategoryId uint) (objets.Categories, error)
}

type TagRepository interface {
	FindTagsByIds(tagIDs []uint) ([]objets.Tags, error)
}
