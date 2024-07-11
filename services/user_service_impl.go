package services

import (
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"Marcketplace/helper"
	"Marcketplace/model/entities"
	"Marcketplace/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		validate:       validate,
	}
}

// Create implements UserService.
func (u *UserServiceImpl) Create(user request.CreateUserRequest) {
	err := u.validate.Struct(user)
	helper.ErrorPanic(err)
	objModel := entities.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	u.userRepository.Save(objModel)
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(userId int) {
	u.userRepository.Delete(userId)
}

// FindAll implements UserService.
func (u *UserServiceImpl) FindAll() []response.UserResponse {
	result := u.userRepository.FindAll()
	var users []response.UserResponse
	for _, value := range result {
		user := response.UserResponse{
			Username: value.Username,
			Email:    value.Email,
		}
		users = append(users, user)
	}
	return users
}

// FindById implements UserService.
func (u *UserServiceImpl) FindById(userId int) response.UserResponse {
	result, err := u.userRepository.FindById(userId)
	helper.ErrorPanic(err)
	user := response.UserResponse{
		Username: result.Username,
		Email:    result.Email,
	}
	return user
}

// Update implements UserService.
func (u *UserServiceImpl) Update(user request.UpdateUserRequest) {
	userData, err := u.userRepository.FindById(int(user.ID))
	helper.ErrorPanic(err)
	userData.Username = user.Username
	userData.Email = user.Email
	userData.Password = user.Password
	u.userRepository.Update(userData)
}

func (u *UserServiceImpl) AuthenticateUser(email, password string) (bool, *entities.User, error) {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return false, nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, nil, err
	}

	return true, user, nil
}
