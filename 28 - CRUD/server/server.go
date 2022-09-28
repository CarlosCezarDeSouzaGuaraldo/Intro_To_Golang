package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Method to INSERT user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Fail to read the request body"))
		return
	}

	var user user

	if err = json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("Fail to convert the body to struct"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Fail to connect in Database"))
		return
	}
	defer db.Close()

	//PREPARE STATEMENT
	statement, err := db.Prepare("INSERT INTO users (name, email) value (?, ?)")
	if err != nil {
		w.Write([]byte("Fail to create the statement"))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Fail to create the statement"))
		return
	}

	id, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Fail to get the inserted ID"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created with success! ID: %d", id)))
}
