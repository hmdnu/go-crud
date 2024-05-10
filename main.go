package main

import (
	"log"
	"net/http"

	"github.com/hmdnubaidillah/go-crud/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/users/{userId}", handlers.GetUser)
	http.HandleFunc("/users/new", handlers.CreateUser)
	http.HandleFunc("/users/del/{userId}", handlers.DeleteUser)
	http.HandleFunc("/users/pat/{userId}", handlers.PatchUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}

}
