package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//CREATE, READ, UPDATE, DELETE

	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", server.RemoveUser).Methods(http.MethodDelete)

	fmt.Println()
	log.Fatal(http.ListenAndServe(":5000", router))
}
