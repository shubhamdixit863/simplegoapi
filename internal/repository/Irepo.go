package repository

import "simplegoapi/internal/entity"

type Repository interface {
	AddData(user entity.User)
	GetData() []entity.User
	UpdateData(user entity.User, id string) entity.User
	DeleteData(name string) entity.User
	GetSingleData(id string) entity.User
}
