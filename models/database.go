package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const username string = "lancy"
const password string = "muschett.26."
const host string = "167.172.187.75"
const port int = 3306
const database string = "vic"

//CreateConnection to the Data Base
func CreateConnection() {
	if connetcion, err := sql.Open("mysql", generateURL()); err != nil {
		panic(err)
	} else {
		db = connetcion
	}
}

//CloseConnection for close de conection to db
func CloseConnection() {
	db.Close()
}
func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
}

//Exec is the wrapper for db.exec to log is an error
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

//Query is the wrapper for db.Query to log is an error
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	return rows, err
}
