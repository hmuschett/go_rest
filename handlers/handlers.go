package handlers

import (
	"log"
	"net/http"
	"rest/models"
	"strconv"

	"github.com/gorilla/mux"
)

//GetUser return an user
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	models.CreateUser()
	user, err := models.GetUser(userID)
	log.Println(err)
	if err != nil {
		models.SendNotFound(w)
		log.Println(err)
	} else {
		models.SendData(w, user)
	}

}
