package repository

import "github.com/go-jwt-test/model"

type UserRepository interface {
	Save(user model.User)
	Update(user model.User)
	Delete(userId string)
	FindById(userId string) (model.User, error)
	FindAll() []model.User
}
