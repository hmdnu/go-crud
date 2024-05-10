package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/hmdnubaidillah/go-crud/data"
	"github.com/hmdnubaidillah/go-crud/models"
	"github.com/hmdnubaidillah/go-crud/utils"
)

func CreateUser(w http.ResponseWriter, req *http.Request) {

	// read req body
	var user models.User

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, "Cant create user", http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	// check if user exist
	for i := range data.Users {
		if user.Name == data.Users[i].Name {
			http.Error(w, "user exist", http.StatusConflict)
			return
		}
	}

	// generate id
	user.ID = utils.GenerateId()

	//
	data.Users = append(data.Users, user)
	fmt.Fprintf(w, "user created")
}

func GetUsers(w http.ResponseWriter, req *http.Request) {

	v, err := json.Marshal(data.Users)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Fprint(w, string(v))

}

func GetUser(w http.ResponseWriter, req *http.Request) {
	userId := req.PathValue("userId")
	uid, err := strconv.Atoi(userId)

	if err != nil {
		log.Fatal(err.Error())
	}

	index := -1

	for i := range data.Users {
		if data.Users[i].ID == uid {
			index = i
		}
	}

	if index == -1 {
		http.Error(w, fmt.Sprintf("cant find userid of %d", uid), http.StatusNotFound)
		return
	}

	v, err := json.Marshal(data.Users[index])

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Fprint(w, string(v))
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	userId := req.PathValue("userId")
	uid, err := strconv.Atoi(userId)

	if err != nil {
		log.Fatal(err.Error())
	}

	// find index that will be deleted
	index := -1

	for i := range data.Users {
		if uid == data.Users[i].ID {
			index = i
		}
	}

	if index == -1 {
		http.Error(w, fmt.Sprintf("cant find user of %d", uid), http.StatusNotFound)
		return
	}

	// reappend
	data.Users = append(data.Users[:index], data.Users[index+1:]...)

	fmt.Fprint(w, "user deleted")
}

func PatchUser(w http.ResponseWriter, req *http.Request) {
	userId := req.PathValue("userId")
	uid, err := strconv.Atoi(userId)

	if err != nil {
		log.Fatal(err.Error())
	}

	// take req body
	var user models.User

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		log.Fatal(err.Error())
		http.Error(w, "cant read input", http.StatusInternalServerError)
		return
	}

	// find matching uid
	index := -1

	for i := range data.Users {
		if data.Users[i].ID == uid {
			index = i
		}
	}

	if index == -1 {
		http.Error(w, fmt.Sprintf("cant find user %d", uid), http.StatusNotFound)
		return
	}

	// edit todo
	data.Users[index].Name = user.Name
	fmt.Fprint(w, "user edited")
}
