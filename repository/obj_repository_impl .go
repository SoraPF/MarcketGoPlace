package repository

import (
	"Marcketplace/helper"
	"Marcketplace/model/objets"
	"errors"
	"log"

	"gorm.io/gorm"
)

type ObjRepositoryImpl struct {
	Db *gorm.DB
}

func NewObjRepositoryImpl(Db *gorm.DB) ObjectRepository {
	return &ObjRepositoryImpl{Db: Db}
}

// Delete implements ObjectRepository.
func (o *ObjRepositoryImpl) Delete(objectId int) {
	var obj objets.Objects
	result := o.Db.Where("id = ? ", objectId).Delete(&obj)
	helper.ErrorPanic(result.Error)
}

// FindAll implements ObjectRepository.
func (o *ObjRepositoryImpl) FindAll() []objets.Objects {
	var obj []objets.Objects
	result := o.Db.Find(&obj)
	helper.ErrorPanic(result.Error)
	return obj
}

// FindById implements ObjectRepository.
func (o *ObjRepositoryImpl) FindById(objectId int) (objets.Objects, error) {
	var obj objets.Objects
	result := o.Db.Find(&objectId)
	if result != nil {
		return obj, nil
	} else {
		return obj, errors.New("note is not found")
	}
}

// Save implements ObjectRepository.
func (o *ObjRepositoryImpl) Save(object objets.Objects) {
	result := o.Db.Create(&object)
	helper.ErrorPanic(result.Error)

}

// Update implements ObjectRepository.
func (o *ObjRepositoryImpl) Update(object objets.Objects) {
	updates := map[string]interface{}{
		"title":       object.Title,
		"price":       object.Price,
		"desc":        object.Desc,
		"status_id":   object.StatusID,
		"category_id": object.CategoryID,
	}
	log.Printf("Updating object with ID: %d, with values: %+v", object.ID, updates)

	// Effectuer la mise à jour
	result := o.Db.Model(&objets.Objects{}).Where("id = ?", object.ID).Updates(updates)

	// Vérifier les erreurs et les lignes affectées
	if result.Error != nil {
		log.Printf("Error updating object: %v", result.Error)
	} else {
		log.Printf("Update successful. Rows affected: %d", result.RowsAffected)
	}
}

func (o *ObjRepositoryImpl) ObjByCategID(CID uint) ([]objets.Objects, error) {
	var obj []objets.Objects
	result := o.Db.Joins("JOIN statuses ON statuses.id = objects.status_id").
		Where("category_id = ? AND statuses.title = ?", CID, "in sale").
		Find(&obj)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("object is not found")
	}
	return obj, nil
}

func (o *ObjRepositoryImpl) ObjByArticleID(CID uint) (objets.Objects, error) {
	var obj objets.Objects
	result := o.Db.Joins("JOIN statuses ON statuses.id = objects.status_id").
		Where("status_id = ? AND statuses.title = ?", CID, "in sale").
		Find(&obj)
	if result.Error != nil {
		return obj, result.Error
	}
	return obj, nil
}

func (o *ObjRepositoryImpl) GetArticles(CID uint, status string) ([]objets.Objects, error) {
	var obj []objets.Objects
	var result *gorm.DB
	if CID == 0 {
		result = o.Db.Joins("JOIN statuses ON statuses.id = objects.status_id").
			Where("statuses.title = ?", status).
			Find(&obj)
	} else {
		result = o.Db.Joins("JOIN statuses ON statuses.id = objects.status_id").
			Where("objects.id = ? AND statuses.title = ?", CID, status).
			Find(&obj)
	}
	if result.Error != nil {
		return obj, result.Error
	}
	return obj, nil
}

func (o *ObjRepositoryImpl) FindByName(name string) (objets.Objects, error) {
	var obj objets.Objects
	var result *gorm.DB
	println(name)
	result = o.Db.Where("objects.title = ?", name).Find(&obj)

	if result.Error != nil {
		return objets.Objects{}, result.Error
	}
	return obj, nil
}
