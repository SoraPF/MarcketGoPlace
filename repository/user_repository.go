package repository

import "Marcketplace/model/entities"

type UserRepository interface {
	Save(user entities.User)
	Update(user entities.User)
	Delete(userId int)
	FindById(userId int) (entities.User, error)
	FindAll() []entities.User
	FindByEmail(email string) (*entities.User, error)
}

type NFARepository interface {
	FindById(NFAId uint) (entities.NFA, error)
}
