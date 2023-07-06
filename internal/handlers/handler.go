package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simplegoapi/internal/dto"
	"simplegoapi/internal/service"
)

type Handler struct {
	Svc service.Service
}

func (hn Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {

	// We can send response from here ----

}

func (hn Handler) AddData(w http.ResponseWriter, r *http.Request) {
	var user dto.AddRequest

	// we will parse the data coming from request into our struct object

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Please send proper data")
	}
	fmt.Println(user)

	hn.Svc.AddData(user)

	fmt.Fprintf(w, "Request Received")

}

func (hn Handler) GetData(w http.ResponseWriter, r *http.Request) {

	users := hn.Svc.GetData()
	fmt.Println(users)

	fmt.Fprintf(w, "Data")

}
