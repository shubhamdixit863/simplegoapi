package repository

import (
	"fmt"
	"gorm.io/gorm"
	"simplegoapi/internal/entity"
)

// Will hold the all the operations of mongodb

type MysqlRepo struct {
	Db *gorm.DB
}

func (repo *MysqlRepo) AddData(user entity.User) {
	result := repo.Db.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		fmt.Println("Error in inserting into Db", result.Error)
		return

	}
	fmt.Println(result.RowsAffected) // That will give us the count of rows inserted

}

// It will get all records

func (repo *MysqlRepo) GetData() []entity.User {

	return nil
}

func (repo *MysqlRepo) UpdateData(user entity.User, id string) entity.User {

	return entity.User{}

}

func (repo *MysqlRepo) DeleteData(id string) entity.User {

	return entity.User{}

}

func (repo *MysqlRepo) GetSingleData(id string) entity.User {

	return entity.User{}

}

// Update One Part

// You can try for filtering the data ,name ,age
//  ---->
