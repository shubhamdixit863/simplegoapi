package repository

import (
	"fmt"
	"simplegoapi/internal/entity"
)

type Repo struct {
	Data []entity.User
}

// We have two different kind of methods
// Pointer Receiver Method
// Value Receiver Method

// Value receiver method

func (repo *Repo) AddData(user entity.User) {
	repo.Data = append(repo.Data, user)

}

func (repo *Repo) GetData() []entity.User {

	return repo.Data
}

func (repo *Repo) UpdateData(user entity.User) entity.User {

	// for updation -->
	for i := 0; i < len(repo.Data); i++ {

		if repo.Data[i].Name == user.Name {

			// We will do the modifications
			repo.Data[i].Age = user.Age
			return user

		}

	}
	return entity.User{}
}

func (repo *Repo) DeleteData(name string) entity.User {

	for i := 0; i < len(repo.Data); i++ {

		if repo.Data[i].Name == name {

			fmt.Println("From service", i)
			// Writing the logic for removing the record
			//	fmt.Println(i)
			//	fmt.Println(repo.Data[:i])
			repo.Data = append(repo.Data[:i], repo.Data[i+1:]...)

			return entity.User{}

		}

	}
	return entity.User{}
}
