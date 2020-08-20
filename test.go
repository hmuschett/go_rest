package main

import (
	"fmt"
	"rest/models"
)

func main() {
	models.CreateConnection()
	user := models.User{Username: "henry2", Password: "1234567"}
	//user.CreateUserMysql()
	err := error(nil)
	user, err = models.GetUser(1)
	fmt.Println(user)
	fmt.Println(err)
	user.Username = "perez"
	models.UpDateUser(user)

	users := models.GetUsers()
	fmt.Println(users)
	models.CloseConnection()
}
