package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest/models"
	"strconv"

	"github.com/gorilla/mux"
)

//GetUser return an user
func GetUser(w http.ResponseWriter, r *http.Request) {
	models.CreateUser()

	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w)
		log.Println(err)
	} else {
		models.SendData(w, user)
	}
}

//CreateUser a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		models.SendData(w, models.SaveUser(user))
	}
}

//UpdateUser an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}

	userResponse := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	models.SendData(w, models.UpDateUser(user.ID, userResponse))
}

//DeleteUser wwerr
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	models.DeleteUser(user.ID)
	models.SendNotContent(w)
}
func getUserByRequest(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	return models.GetUser(userID)
}
