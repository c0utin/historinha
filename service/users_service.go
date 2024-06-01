package service

import (
	"github.com/c0utin/historinha/data/request"
	"github.com/c0utin/historinha/data/response"
)

type UsersService interface {
	Create(users request.CreateUsersRequest)
	Update(users request.UpdateUsersRequest)
	Delete(usersId int)
	FindById(userId int) response.UsersResponse
	FindAll() []response.UsersResponse
}
