package service

import (
	"github.com/go-jwt-test/data/request"
	"github.com/go-jwt-test/data/response"
)

type UserService interface {
	Create(user request.CreateUserRequest)
	Update(user request.UpdateUserRequest)
	Delete(userId string)
	FindByid(userId string) response.UserResponse
	FindAll() []response.UserResponse
}
