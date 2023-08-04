package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"simplegoapi/internal/dto"
	"simplegoapi/internal/service"
)

type Handler struct {
	Svc service.IService
}

func NewHandler(Svc service.IService) Handler {
	return Handler{
		Svc: Svc,
	}

}

func (hn Handler) AddData(w http.ResponseWriter, r *http.Request) {
	var user dto.AddRequest

	// we will parse the data coming from request into our struct object

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Please send proper data")
	}

	hn.Svc.AddData(user) // We can mock this behaviour ---->

	fmt.Fprintf(w, "Request Received")

}

func (hn Handler) GetData(w http.ResponseWriter, r *http.Request) {

	users, err := hn.Svc.GetData()

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	//fmt.Println(users)

	// Two ways go provides to convert a struct into a json string
	// First in json.Marshal
	// Second is json.Encode

	marshal, err := json.Marshal(users)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(marshal))

}

func (hn Handler) UpdateData(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	var user dto.AddRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Please send proper data")
	}

	err = hn.Svc.UpdateData(user, id)
	if err != nil {
		fmt.Fprintf(w, err.Error())

		return
	}

	fmt.Fprintf(w, "User Updated")
}

func (hn Handler) DeleteData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := hn.Svc.DeleteData(id)
	if err != nil {
		fmt.Fprintf(w, err.Error())

		return
	}

	fmt.Fprintf(w, "User Deleted")
}

func (hn Handler) GetSingleData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	dataFromService, err := hn.Svc.GetSingleData(id)
	if err != nil {
		fmt.Fprintf(w, err.Error())

	}

	// We have to marshal it into a json string

	marshal, err := json.Marshal(dataFromService)
	if err != nil {
		fmt.Fprintf(w, "Error in marshalling")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(marshal))
}
