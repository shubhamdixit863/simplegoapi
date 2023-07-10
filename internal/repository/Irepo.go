package repository

import "simplegoapi/internal/entity"

type Repository interface {
	AddData(user entity.User)
	GetData() []entity.User
	UpdateData(user entity.User) entity.User
	DeleteData(name string) entity.User
}
