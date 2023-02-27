package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/variations", variations).Methods("GET")
	router.HandleFunc("/edit", edit).Methods("GET")
	log.Fatal(http.ListenAndServe(":9090", router))

}
