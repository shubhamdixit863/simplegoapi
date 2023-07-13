package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"simplegoapi/internal/entity"
)

// Will hold the all the operations of mongodb

type MongoRepo struct {
	Db *mongo.Database
}

func (repo *MongoRepo) AddData(user entity.User) {

	fmt.Println(user)

	_, err := repo.Db.Collection("user").InsertOne(context.Background(), user)
	if err != nil {
		return
	}
	//fmt.Println(one)
}

// It will get all records

func (repo *MongoRepo) GetData() []entity.User {

	var slc []entity.User

	//filter := bson.D{{"name", "ram"}}
	filter := bson.D{}
	find, err := repo.Db.Collection("user").Find(context.Background(), filter)
	if err != nil {
		return nil
	}

	// We have to iterate over this find to get the data

	// find.Next will keep returning the boolean value until we have the data in our find
	// or until we consume all the records
	for find.Next(context.TODO()) {
		d := entity.User{}
		err := find.Decode(&d)
		if err != nil {
			return nil
		}

		// we will append the records here

		slc = append(slc, d)

	}

	return slc

}

func (repo *MongoRepo) UpdateData(user entity.User, id string) entity.User {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.User{}
	}

	// bson.M is for unstructured data ,used as map
	//filter := bson.M{"name": bson.M{"$regex": user.Name, "$options": "i"}}  // regex search is supported by mongodb ,i,,g global search
	// bson.D is a combination of slices
	//options := options.Update().SetUpsert(true) // upsert means update if available otheriwse insert
	update := bson.D{{"$set", bson.D{

		{"name", user.Name},
		{"age", user.Age},
	}}}
	filter := bson.D{{"_id", objectId}}
	_, err = repo.Db.Collection("user").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return entity.User{}
	}

	return user

}

func (repo *MongoRepo) DeleteData(id string) entity.User {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.User{}
	}
	filter := bson.D{{"_id", objectId}}
	repo.Db.Collection("user").DeleteOne(context.Background(), filter)
	return entity.User{}

}

func (repo *MongoRepo) GetSingleData(id string) entity.User {

	// OF getting single data from the db

	d := entity.User{}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.User{}
	}
	filter := bson.D{{"_id", objectId}}

	err = repo.Db.Collection("user").FindOne(context.Background(), filter).Decode(&d)

	fmt.Println(d)
	if err != nil {
		return entity.User{}
	}

	return d

}

// Update One Part

// You can try for filtering the data ,name ,age
//  ---->
