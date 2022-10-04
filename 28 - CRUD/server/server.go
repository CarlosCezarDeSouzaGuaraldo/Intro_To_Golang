package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
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
		w.Write([]byte("Fail to create user"))
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

// Method to GET all users in database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Fail to connect in Database"))
		return
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM users")
	if err != nil {
		w.Write([]byte("Fail to get users"))
		return
	}
	defer data.Close()

	var users []user
	for data.Next() {
		var user user

		if err := data.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Fail to scan users"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Fail to convert users to JSON"))
		return
	}
}

// Method to GET a unique user in database
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Fail to convert param to integer"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Fail to connect in Database"))
		return
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM users WHERE id = ?", ID)
	if err != nil {
		w.Write([]byte("Fail to get user"))
		return
	}
	defer data.Close()

	var user user

	if data.Next() {
		if err := data.Scan(user.ID, user.Name, user.Email); err != nil {
			w.Write([]byte("Fail to scan user"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Fail to convert users to JSON"))
		return
	}
}

// Method to UPDATE a unique user in database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Fail to convert param to integer"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Fail to read the request body"))
		return
	}

	var user user
	if err := json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("Fail to convert the body to struct"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Fail to connect in Database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		w.Write([]byte("Fail to create the statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Fail to update user"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Method to UPDATE a unique user in database
func RemoveUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Fail to convert param to integer"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Fail to connect in Database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		w.Write([]byte("Fail to create the statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Fail to delete user"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
