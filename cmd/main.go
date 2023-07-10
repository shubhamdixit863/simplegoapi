package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"simplegoapi/internal/handlers"
	"simplegoapi/internal/repository"
	"simplegoapi/internal/service"
)

func Mongoconnect(uri string) *mongo.Database {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected with mongodb")
	return client.Database("mong_tute")
}

func main() {

	r := mux.NewRouter()

	mongodb := Mongoconnect("mongodb://localhost:27017")

	// db object
	// in memory db object
	//mp := make(map[string]string)
	//mp := []entity.User{}

	// Repository objects
	// inmemory repo
	//repo := repository.Repo{Data: mp}
	// mongodb repo
	mongodRepo := repository.MongoRepo{Db: mongodb}

	// Service object

	svc := service.Service{Repository: &mongodRepo}

	// handler object

	hn := handlers.Handler{Svc: svc}

	r.HandleFunc("/", hn.HomeHandler)
	r.HandleFunc("/user", hn.AddData).Methods("POST")
	r.HandleFunc("/user", hn.GetData).Methods("GET")
	r.HandleFunc("/user", hn.UpdateData).Methods("PUT")
	r.HandleFunc("/user/{name}", hn.DeleteData).Methods("DELETE")

	log.Fatalln(http.ListenAndServe(":8080", r))

}
