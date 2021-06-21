package user

import "server-api/entity"

type Service interface {
	SaveNewUser(input entity.UserInput) (entity.User, error)
	GetUserByID()
	EditUserByID()
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}
