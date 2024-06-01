package service

import (
	"github.com/c0utin/historinha/data/request"
	"github.com/c0utin/historinha/data/response"
	"github.com/c0utin/historinha/helper"
	"github.com/c0utin/historinha/model"
	"github.com/c0utin/historinha/repository"

	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewUsersServiceImpl(userRepository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: userRepository,
		Validate:        validate,
	}
}

// Create implements UsersService
func (u *UsersServiceImpl) Create(users request.CreateUsersRequest) {
	err := u.Validate.Struct(users)
	helper.ErrorPanic(err)
	userModel := model.Users{
		Name: users.Name,
	}
	u.UsersRepository.Save(userModel)
}

// Delete implements UsersService
func (u *UsersServiceImpl) Delete(usersId int) {
	u.UsersRepository.Delete(usersId)
}

// FindAll implements UsersService
func (u *UsersServiceImpl) FindAll() []response.UsersResponse {
	result := u.UsersRepository.FindAll()

	var users []response.UsersResponse
	for _, value := range result {
		user := response.UsersResponse{
			ID:   value.ID,
			Name: value.Name,
		}
		users = append(users, user)
	}

	return users
}

// FindById implements UsersService
func (u *UsersServiceImpl) FindById(usersId int) response.UsersResponse {
	userData, err := u.UsersRepository.FindById(usersId)
	helper.ErrorPanic(err)

	userResponse := response.UsersResponse{
		ID:   userData.ID,
		Name: userData.Name,
	}
	return userResponse
}

// Update implements UsersService
func (u *UsersServiceImpl) Update(users request.UpdateUsersRequest) {
	userData, err := u.UsersRepository.FindById(users.ID)
	helper.ErrorPanic(err)
	userData.Name = users.Name
	u.UsersRepository.Update(userData)
}

