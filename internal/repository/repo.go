package repository

import (
	"fmt"
	"simplegoapi/internal/entity/mongodb"
)

type Repo struct {
	Data []mongodb.User
}

// We have two different kind of methods
// Pointer Receiver Method
// Value Receiver Method

// Value receiver method

func (repo *Repo) AddData(user mongodb.User) {
	repo.Data = append(repo.Data, user)

}

func (repo *Repo) GetData() []mongodb.User {

	return repo.Data
}

func (repo *Repo) UpdateData(user mongodb.User) mongodb.User {

	// for updation -->
	for i := 0; i < len(repo.Data); i++ {

		if repo.Data[i].Name == user.Name {

			// We will do the modifications
			repo.Data[i].Age = user.Age
			return user

		}

	}
	return mongodb.User{}
}

func (repo *Repo) DeleteData(name string) mongodb.User {

	for i := 0; i < len(repo.Data); i++ {

		if repo.Data[i].Name == name {

			fmt.Println("From service", i)
			// Writing the logic for removing the record
			//	fmt.Println(i)
			//	fmt.Println(repo.Data[:i])
			repo.Data = append(repo.Data[:i], repo.Data[i+1:]...)

			return mongodb.User{}

		}

	}
	return mongodb.User{}
}
