package services

import (
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"Marcketplace/model/entities"
)

type UserService interface {
	Create(user request.CreateUserRequest)
	Update(user request.UpdateUserRequest)
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindAll() []response.UserResponse
	AuthenticateUser(email, password string) (bool, *entities.User, error)
	FindUser(userId int) *entities.User
}
