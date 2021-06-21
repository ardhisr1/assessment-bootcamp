package user

import (
	"errors"
	"fmt"
	"server-api/entity"
)

type Service interface {
	SaveNewUser(input entity.UserInput) (entity.User, error)
	GetUserByID(UserID string) (entity.User, error)
	EditUserByID(UserID string, data entity.UserUpdateInput) (entity.User, error)
	Login(account entity.UserLogin)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) SaveNewUser(input entity.UserInput) (entity.User, error) {
	// encrypt, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	// if err != nil {
	// 	return entity.User{}, err
	// }

	newUser := entity.User{
		FullName: input.FullName,
		Address:  input.Address,
		Email:    input.Email,
		Password: input.Password,
	}

	createdUser, err := s.repository.Create(newUser)
	if err != nil {
		return createdUser, err
	}
	return createdUser, nil
}

func (s *service) GetUserByID(UserID string) (entity.User, error) {
	user, err := s.repository.FindByID(UserID)
	if err != nil {
		return entity.User{}, err
	}

	if user.ID == 0 {
		NewError := fmt.Sprintf("user id %s not found", UserID)
		return entity.User{}, errors.New(NewError)
	}

	return user, nil
}

func (s *service) Login(account entity.UserLogin) (entity.User, error) {
	user, err := s.repository.FindByEmai(account.Email)
	if err != nil {
		return entity.User{}, err
	}

	if user.Password != account.Password {
		return user, errors.New("password invalid")
	}
	// if err := bycrypt.CompareHashAndPassword([]byte(user.Password), []byte(account.Password)); err != nil {
	// 	return user, errors.New("password invalid")
	// }

	return user, nil
}
