package repository

import (
	"Marcketplace/data/request"
	"Marcketplace/helper"
	"Marcketplace/model/entities"
	"errors"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewuserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Delete implements userRepository.
func (u *UserRepositoryImpl) Delete(userId int) {
	var user entities.User
	result := u.Db.Where("id = ? ", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// FindAll implements userRepository.
func (u *UserRepositoryImpl) FindAll() []entities.User {
	var user []entities.User
	result := u.Db.Find(&user)
	helper.ErrorPanic(result.Error)
	return user
}

// FindById implements userRepository.
func (u *UserRepositoryImpl) FindById(userId int) (entities.User, error) {
	var obj entities.User
	result := u.Db.Find(&userId)
	if result != nil {
		return obj, nil
	} else {
		return obj, errors.New("note is not found")
	}
}

// Save implements userRepository.
func (u *UserRepositoryImpl) Save(user entities.User) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

// Update implements userRepository.
func (u *UserRepositoryImpl) Update(user entities.User) {
	var updateUser = request.UpdateUserRequest{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	result := u.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
