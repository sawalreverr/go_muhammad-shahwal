package user

import (
	"errors"
	"go-wishlist-api/utils/helper"
)

type Service interface {
	Register(userInput UserRegister) (User, error)
	Login(userInput UserLogin) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(userInput UserRegister) (User, error) {
	hashesPassword := helper.MD5Hash(userInput.Password)

	user := User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: hashesPassword,
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(userInput UserLogin) (User, error) {
	userFound, err := s.repository.FindByEmail(userInput.Email)
	if err != nil {
		return userFound, err
	}

	hashesPassword := helper.MD5Hash(userInput.Password)
	if userFound.Password != hashesPassword {
		return userFound, errors.New("invalid password")
	}

	return userFound, nil
}

func (s *service) GetUserByEmail(email string) (User, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Email == email {
		return user, errors.New("email already exists")
	}

	return user, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on with that ID")
	}

	return user, nil
}
