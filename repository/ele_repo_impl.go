package repository

import (
	"Marcketplace/helper"
	"Marcketplace/model/objets"

	"gorm.io/gorm"
)

type ElemRepositoryImpl struct {
	Db *gorm.DB
}

func NewElemRepositoryImpl(Db *gorm.DB) ElemRepository {
	return &ElemRepositoryImpl{Db: Db}
}

// FindAllCategories implements ElemRepository.
func (e *ElemRepositoryImpl) FindAllCategories() []objets.Categories {
	var ele []objets.Categories
	result := e.Db.Find(&ele)
	helper.ErrorPanic(result.Error)
	return ele
}

// FindAllTags implements ElemRepository.
func (e *ElemRepositoryImpl) FindAllTags() []objets.Tags {
	var ele []objets.Tags
	result := e.Db.Find(&ele)
	helper.ErrorPanic(result.Error)
	return ele
}
