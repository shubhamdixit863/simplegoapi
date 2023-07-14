package repository

import (
	"fmt"
	"gorm.io/gorm"
	"simplegoapi/internal/entity/mysql"
	"strconv"
)

// Will hold the all the operations of mongodb

type MysqlRepo struct {
	Db *gorm.DB
}

func (repo *MysqlRepo) AddData(user mysql.User) error {
	result := repo.Db.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		fmt.Println("Error in inserting into Db", result.Error)
		return result.Error

	}
	return nil
}

// It will get all records

func (repo *MysqlRepo) GetData() ([]mysql.User, error) {
	var data []mysql.User
	result := repo.Db.Find(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func (repo *MysqlRepo) UpdateData(user mysql.User, id string) (mysql.User, error) {

	// We will convert the id here to number or to int
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return mysql.User{}, err
	}
	// If save value does not contain primary key, it will execute Create, otherwise it will execute Update (with all fields).
	user.Id = atoi

	result := repo.Db.Save(&user)

	if result.Error != nil {

		return mysql.User{}, result.Error
	}
	return user, nil

}

func (repo *MysqlRepo) DeleteData(id string) (mysql.User, error) {

	// We will convert the id here to number or to int
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return mysql.User{}, err
	}
	var user mysql.User
	user.Id = atoi
	result := repo.Db.Delete(&user)
	if result.Error != nil {
		return mysql.User{}, result.Error
	}

	return mysql.User{}, nil

}

func (repo *MysqlRepo) GetSingleData(id string) (mysql.User, error) {

	// We will convert the id here to number or to int
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return mysql.User{}, err
	}

	var data mysql.User
	result := repo.Db.First(&data, atoi)
	if result.Error != nil { //null

		return data, err

	}

	return data, nil

}

// Update One Part

// You can try for filtering the data ,name ,age
//  ---->
