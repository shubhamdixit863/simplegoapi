package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"simplegoapi/internal/handlers"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/user", handlers.HomeHandler)

	log.Fatalln(http.ListenAndServe(":8080", r))

}
