package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name  string
	Email string
}

type CreateUserResponse struct {
	Message string
	User    User
}

type ReadUserResponse struct {
	User User
}

var users []User

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var user User
	json.Unmarshal(body, &user)

	if user.Name == "" || user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Name and email are required")
		return
	}

	users = append(users, user)

	response := CreateUserResponse{
		Message: "User created successfully",
		User:    user,
	}

	responseBody, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func ReadUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	email := query.Get("email")

	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Email is required")
		return
	}

	for _, user := range users {
		if user.Email == email {
			response := ReadUserResponse{
				User: user,
			}

			responseBody, _ := json.Marshal(response)

			w.Header().Set("Content-Type", "application/json")
			w.Write(responseBody)
			return
		}
	}

	fmt.Fprintf(w, "User not found")
}

func main() {
	http.HandleFunc("/createUser", CreateUserHandler)
	http.HandleFunc("/readUser", ReadUserHandler)

	http.ListenAndServe(":8888", nil)
}
