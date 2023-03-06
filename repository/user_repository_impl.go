package repository

import (
	"errors"

	"github.com/go-jwt-test/data/request"
	"github.com/go-jwt-test/helper"
	"github.com/go-jwt-test/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(w *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: w}

}

// Delete implements UserRepository
func (c *UserRepositoryImpl) Delete(userId string) {
	var user model.User
	result := c.Db.Where("uid = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UserRepository
func (c *UserRepositoryImpl) FindAll() []model.User {
	var user []model.User
	result := c.Db.Find(&user)
	helper.ErrorPanic(result.Error)
	return user
}

// FindById implements UserRepository
func (c *UserRepositoryImpl) FindById(userId string) (model.User, error) {
	var user model.User
	result := c.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// Save implements UserRepository
func (c *UserRepositoryImpl) Save(user model.User) {
	result := c.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

// Update implements UserRepository
func (c *UserRepositoryImpl) Update(user model.User) {
	var updateUser = request.UpdateUserRequest{
		Uid:      user.Uid.String(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	result := c.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}
