package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func MysqlConnect(uri string) *gorm.DB {

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected with Mysql")

	return db
}

func main() {

	r := mux.NewRouter()

	//mongodb := Mongoconnect("mongodb://localhost:27017")

	dsn := "admin:o9Uusjfn@tcp(mysql-135552-0.cloudclusters.net:17741)/mysqlGolang?charset=utf8mb4&parseTime=True&loc=Local"
	db := MysqlConnect(dsn)

	// db object
	// in memory db object
	//mp := make(map[string]string)
	//mp := []entity.User{}

	// Repository objects
	// inmemory repo
	//repo := repository.Repo{Data: mp}
	// mongodb repo
	//	mongodRepo := repository.MongoRepo{Db: mongodb}

	// mysql Rep
	mysqlRepo := repository.MysqlRepo{Db: db}

	// Service object

	svc := service.Service{Repository: &mysqlRepo}

	// handler object

	hn := handlers.Handler{Svc: svc}

	//r.HandleFunc("/", hn.HomeHandler)
	r.HandleFunc("/user", hn.AddData).Methods("POST")
	r.HandleFunc("/user", hn.GetData).Methods("GET")
	r.HandleFunc("/user/{id}", hn.GetSingleData).Methods("GET")
	r.HandleFunc("/user/{id}", hn.UpdateData).Methods("PUT")
	r.HandleFunc("/user/{id}", hn.DeleteData).Methods("DELETE")

	log.Fatalln(http.ListenAndServe(":8080", r))

}
