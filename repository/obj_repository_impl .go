package repository

import (
	"Marcketplace/data/request"
	"Marcketplace/helper"
	"Marcketplace/model/objets"
	"errors"

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
	var tagIDs []uint
	for _, tag := range object.Tags {
		tagIDs = append(tagIDs, tag.ID)
	}
	var updateObj = request.UpdateObjRequest{
		ID:         object.ID,
		Title:      object.Title,
		Price:      object.Price,
		Desc:       object.Desc,
		StatusID:   object.StatusID,
		CategoryID: object.CategoryID,
		Tags:       tagIDs,
	}
	result := o.Db.Model(&object).Updates(updateObj)
	helper.ErrorPanic(result.Error)
}

func (o *ObjRepositoryImpl) ObjByCategID(CID uint) ([]objets.Objects, error) {
	var obj []objets.Objects
	result := o.Db.Where("category_id = ?", CID).Find(&obj)
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
	result := o.Db.Where("id = ?", CID).Find(&obj)
	if result.Error != nil {
		return obj, result.Error
	}
	return obj, nil
}
