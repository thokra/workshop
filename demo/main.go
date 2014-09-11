package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

var users []User

func init() {
	users = []User{
		User{"yellowlion265", "Caleb Watson"},
		User{"crazypeacock992", "Ivan Powell"},
		User{"smalllion550", "Corey Robertson"},
		User{"tinypeacock741", "Ashley Wade"},
	}
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", allUsers).Methods("GET")

	http.ListenAndServe(":3000", r)
}
