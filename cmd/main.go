package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"simplegoapi/internal/entity"
	"simplegoapi/internal/handlers"
	"simplegoapi/internal/repository"
	"simplegoapi/internal/service"
)

func main() {

	r := mux.NewRouter()

	// db object
	// in memory db object
	//mp := make(map[string]string)
	mp := []entity.User{}

	// Repository object

	repo := repository.Repo{Data: mp}

	// Service object

	svc := service.Service{Repository: &repo}

	// handler object

	hn := handlers.Handler{Svc: svc}

	r.HandleFunc("/", hn.HomeHandler)
	r.HandleFunc("/user", hn.AddData).Methods("POST")
	r.HandleFunc("/user", hn.GetData).Methods("GET")
	r.HandleFunc("/user", hn.UpdateData).Methods("PUT")
	r.HandleFunc("/user/{name}", hn.DeleteData).Methods("DELETE")

	log.Fatalln(http.ListenAndServe(":8080", r))

}
