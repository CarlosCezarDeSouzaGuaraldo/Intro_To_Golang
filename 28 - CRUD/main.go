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

	fmt.Println()
	log.Fatal(http.ListenAndServe(":5000", router))
}
