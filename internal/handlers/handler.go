package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	// We can send response from here ----

	fmt.Fprintf(w, "Hello world This is my first Route")

}
