package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var users = []User{}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	var payload User

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		FirstName: "Harshwardhan",
		LastName:  "Chandanshive",
	}

	users = append(users, u)

	w.WriteHeader(http.StatusCreated)

}
