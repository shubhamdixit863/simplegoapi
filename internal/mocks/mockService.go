package mocks

import (
	"github.com/stretchr/testify/mock"
	"simplegoapi/internal/dto"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) AddData(user dto.AddRequest) error {
	m.Called(user) // reflection and we are imitating the call
	return nil
}

func (m *MockService) GetData() ([]dto.UserResponseDto, error) {
	// reflection and we are imitating the call
	return nil, nil
}

func (m *MockService) UpdateData(user dto.AddRequest, id string) error {
	args := m.Called(user) // reflection and we are imitating the call
	return args.Error(1)
}

func (m *MockService) GetSingleData(id string) (dto.AddRequest, error) {
	args := m.Called(id) // reflection and we are imitating the call
	return args.Get(0).(dto.AddRequest), args.Error(1)
}

func (m *MockService) DeleteData(id string) error {
	args := m.Called(id) // reflection and we are imitating the call
	return args.Error(1)
}
