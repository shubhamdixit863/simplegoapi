package repository

import (
	"simplegoapi/internal/entity/mongodb"
)

type Repository interface {
	AddData(user mongodb.User)
	GetData() []mongodb.User
	UpdateData(user mongodb.User, id string) mongodb.User
	DeleteData(name string) mongodb.User
	GetSingleData(id string) mongodb.User
}
