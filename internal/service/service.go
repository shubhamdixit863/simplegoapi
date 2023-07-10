package service

import (
	"fmt"
	"simplegoapi/internal/dto"
	"simplegoapi/internal/entity"
	"simplegoapi/internal/repository"
)

type Service struct {
	Repository repository.Repository
}

func (svc Service) AddData(user dto.AddRequest) {

	// We will do some operation on our repository

	userEntity := entity.User{
		Name: user.Name,
		Age:  user.Age,
	}

	fmt.Println(userEntity)

	svc.Repository.AddData(userEntity)

}

func (svc Service) GetData() []dto.AddRequest {

	users := []dto.AddRequest{}
	entityUser := svc.Repository.GetData()
	for _, value := range entityUser {

		fmt.Println(value)

		userDto := dto.AddRequest{
			Name: value.Name,
			Age:  value.Age,
		}
		users = append(users, userDto)

	}

	return users

}

func (svc Service) UpdateData(user dto.AddRequest) {

	userEntity := entity.User{
		Name: user.Name,
		Age:  user.Age,
	}

	// We will call our repo method to update the record

	svc.Repository.UpdateData(userEntity)
}

func (svc Service) DeleteData(name string) {

	svc.Repository.DeleteData(name)

}
