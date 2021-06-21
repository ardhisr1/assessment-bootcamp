package user

import (
	"errors"
	"fmt"
	"server-api/entity"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	SaveNewUser(input entity.UserInput) (UserFormat, error)
	GetUserByID(UserID string) (UserFormat, error)
	EditUserByID(UserID string, data entity.UserUpdateInput) (UserFormat, error)
	Login(account entity.UserLogin) (UserFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) SaveNewUser(input entity.UserInput) (UserFormat, error) {
	encrypt, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return UserFormat{}, err
	}

	newUser := entity.User{
		FullName: input.FullName,
		Address:  input.Address,
		Email:    input.Email,
		Password: string(encrypt),
	}

	createdUser, err := s.repository.Create(newUser)
	if err != nil {
		return UserFormat{}, err
	}

	formatedUser := FormatUser(createdUser.ID, createdUser.FullName, createdUser.Address, createdUser.Email)
	return formatedUser, nil
}

func (s *service) GetUserByID(UserID string) (UserFormat, error) {
	user, err := s.repository.FindByID(UserID)
	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		NewError := fmt.Sprintf("user id %s not found", UserID)
		return UserFormat{}, errors.New(NewError)
	}

	formatedUser := FormatUser(user.ID, user.FullName, user.Address, user.Email)
	return formatedUser, nil
}

func (s *service) Login(account entity.UserLogin) (UserFormat, error) {
	user, err := s.repository.FindByEmai(account.Email)
	if err != nil {
		return UserFormat{}, err
	}

	// if user.Password != account.Password {
	// 	return user, errors.New("password invalid")
	// }
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(account.Password)); err != nil {
		return UserFormat{}, errors.New("password invalid")
	}

	formatedUser := FormatUser(user.ID, user.FullName, user.Address, user.Email)
	return formatedUser, nil
}

func (s *service) EditUserByID(UserID string, data entity.UserUpdateInput) (UserFormat, error) {
	var user entity.User

	formatedUser := FormatUser(user.ID, user.FullName, user.Address, user.Email)
	return formatedUser, nil
}
