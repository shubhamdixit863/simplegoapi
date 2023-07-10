package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"simplegoapi/internal/entity"
)

// Will hold the all the operations of mongodb

type MongoRepo struct {
	Db *mongo.Database
}

func (repo *MongoRepo) AddData(user entity.User) {

	one, err := repo.Db.Collection("user").InsertOne(context.Background(), bson.D{
		{"name", user.Name},
		{"age", user.Age},
	})
	if err != nil {
		return
	}
	fmt.Println(one)
}

func (repo *MongoRepo) GetData() []entity.User {

	return nil

}

func (repo *MongoRepo) UpdateData(user entity.User) entity.User {

	return user

}

func (repo *MongoRepo) DeleteData(name string) entity.User {
	return entity.User{}

}
