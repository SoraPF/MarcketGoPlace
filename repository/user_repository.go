package repository

import "Marcketplace/model/entities"

type UserRepository interface {
	Save(user entities.User)
	Update(user entities.User)
	Delete(userId int)
	FindById(userId uint) (entities.User, error)
	FindAll() []entities.User
	FindByEmail(email string) (*entities.User, error)
	CreateNFA(nfa *entities.NFA) error
	FindNFA(nfa *uint) (*entities.NFA, error)
}
