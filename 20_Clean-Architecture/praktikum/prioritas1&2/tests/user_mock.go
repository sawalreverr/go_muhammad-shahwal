package tests

import (
	"go-wishlist-api/user"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUserByEmail(email string) (user.User, error) {
	args := m.Called(email)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockUserService) GetUserByID(id int) (user.User, error) {
	args := m.Called(id)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockUserService) Register(input user.UserRegister) (user.User, error) {
	args := m.Called(input)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockUserService) Login(input user.UserLogin) (user.User, error) {
	args := m.Called(input)
	return args.Get(0).(user.User), args.Error(1)
}
