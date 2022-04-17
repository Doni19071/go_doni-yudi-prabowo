package service

import "belajar-go-echo/model"

type MockUserService struct {
	data []model.User
}

func (ps *MockUserService) Add(user model.User) (model.User, error) {
	ps.data = append(ps.data, user)
	return user, nil
}

func (ps *MockUserService) Get() ([]model.User, error) {
	return ps.data, nil
}

func NewMockUserService() *MockUserService {
	return &MockUserService{
		data: []model.User{},
	}
}
