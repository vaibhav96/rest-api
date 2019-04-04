package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// our main function
func main() {
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

// HandleError conveniently handles error.
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}


