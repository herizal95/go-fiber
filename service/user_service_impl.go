package service

import (
	"github.com/go-jwt-test/data/request"
	"github.com/go-jwt-test/data/response"
	"github.com/go-jwt-test/helper"
	"github.com/go-jwt-test/model"
	"github.com/go-jwt-test/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

// Create implements UserService
func (c *UserServiceImpl) Create(user request.CreateUserRequest) {
	err := c.validate.Struct(user)
	helper.ErrorPanic(err)
	userModel := model.User{
		Uid:      uuid.New(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	c.userRepository.Save(userModel)

}

// Delete implements UserService
func (c *UserServiceImpl) Delete(userId string) {
	c.userRepository.Delete(userId)
}

// FindAll implements UserService
func (c *UserServiceImpl) FindAll() []response.UserResponse {
	result := c.userRepository.FindAll()
	var user []response.UserResponse

	for _, value := range result {
		users := response.UserResponse{
			Uid:      value.Uid,
			Name:     value.Name,
			Email:    value.Email,
			Password: value.Password,
		}
		user = append(user, users)
	}
	return user
}

// FindByid implements UserService
func (c *UserServiceImpl) FindByid(userId string) response.UserResponse {
	userData, err := c.userRepository.FindById(userId)
	helper.ErrorPanic(err)
	userResponse := response.UserResponse{
		Uid:      userData.Uid,
		Name:     userData.Name,
		Email:    userData.Email,
		Password: userData.Password,
	}
	return userResponse
}

// Update implements UserService
func (c *UserServiceImpl) Update(user request.UpdateUserRequest) {
	userData, err := c.userRepository.FindById(user.Uid)
	helper.ErrorPanic(err)
	userData.Name = user.Name
	userData.Email = user.Email
	userData.Password = user.Password
	c.userRepository.Update(userData)
}
