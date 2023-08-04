package service

import (
	"fmt"
	"simplegoapi/internal/dto"
	"simplegoapi/internal/entity/mysql"
	"simplegoapi/internal/repository"
	"strings"
)

type IService interface {
	AddData(user dto.AddRequest) error
	GetData() ([]dto.UserResponseDto, error)
	UpdateData(user dto.AddRequest, id string) error
	GetSingleData(id string) (dto.AddRequest, error)
	DeleteData(id string) error
}

type Service struct {
	Repository repository.RelationalDbRepository
}

func (svc Service) AddData(user dto.AddRequest) error {

	// We will do some operation on our repository

	userEntity := mysql.User{
		Name: user.Name,
		Age:  user.Age,
	}

	fmt.Println(userEntity)

	return svc.Repository.AddData(userEntity)

}

func (svc Service) GetData() ([]dto.UserResponseDto, error) {

	users := []dto.UserResponseDto{}
	entityUser, err := svc.Repository.GetData()

	if err != nil {
		return nil, err
	}
	for _, value := range entityUser {

		fmt.Println(value)

		userDto := dto.UserResponseDto{
			Id:   value.Id,
			Name: strings.ToUpper(value.Name),
			Age:  value.Age,
		}
		users = append(users, userDto)

	}

	return users, nil

}

func (svc Service) UpdateData(user dto.AddRequest, id string) error {

	userEntity := mysql.User{
		Name: user.Name,
		Age:  user.Age,
	}

	// We will call our repo method to update the record

	_, err := svc.Repository.UpdateData(userEntity, id)
	if err != nil {
		return err
	}
	return nil
}

func (svc Service) DeleteData(id string) error {

	_, err := svc.Repository.DeleteData(id)
	if err != nil {
		return err
	}

	return nil

}

func (svc Service) GetSingleData(id string) (dto.AddRequest, error) {

	d, err := svc.Repository.GetSingleData(id)
	if err != nil {
		return dto.AddRequest{}, err
	}

	return dto.AddRequest{
		Name: d.Name,
		Age:  d.Age,
	}, nil

}
