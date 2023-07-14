package repository

import (
	"simplegoapi/internal/entity/mysql"
)

// For relational databases

type RelationalDbRepository interface {
	AddData(user mysql.User) error
	GetData() ([]mysql.User, error)
	UpdateData(user mysql.User, id string) (mysql.User, error)
	DeleteData(name string) (mysql.User, error)
	GetSingleData(id string) (mysql.User, error)
}
