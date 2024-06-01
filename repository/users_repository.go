package repository 

import "github.com/c0utin/historinha/model"

type UsersRepository interface {
	Save(users model.Users)
	Update(users model.Users)
	Delete(userId int)
	FindById(userId int) (users model.Users, err error)
	FindAll() []model.Users
}
