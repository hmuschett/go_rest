package models

import "errors"

//User struct represent an user
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[int]User)

func CreateUser() {
	user := User{ID: 1, Username: "Henry", Password: "qwer"}
	users[user.ID] = user
}

func GetUser(ID int) (User, error) {
	if user, err := users[ID]; err {
		return user, nil
	}
	return User{}, errors.New("El usuario no existe")
}
