package repository

import (
	"errors"
	"github.com/c0utin/historinha/data/request"
	"github.com/c0utin/historinha/helper"
	"github.com/c0utin/historinha/model"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

// Delete implements UsersRepository
func (u *UsersRepositoryImpl) Delete(userId int) {
	var user model.Users
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UsersRepository
func (u *UsersRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	result := u.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindById implements UsersRepository
func (u *UsersRepositoryImpl) FindById(userId int) (users model.Users, err error) {
	var user model.Users
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// Save implements UsersRepository
func (u *UsersRepositoryImpl) Save(user model.Users) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

// Update implements UsersRepository
func (u *UsersRepositoryImpl) Update(users model.Users) {
	var updateUser = request.UpdateUsersRequest{
		ID:   users.ID,
		Name: users.Name,
	}
	result := u.Db.Model(&users).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

