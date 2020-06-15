package main

import (
	"log"
	"net/http"

	"rest/handlers"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/api/v1/users", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")

	log.Println("the server is lisening on 3000 port")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
