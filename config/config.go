package config

import (
	"database/sql"
	"fmt"
	"learn/REST/utils"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username string = utils.GetEnvironment("username")
	password string = utils.GetEnvironment("password")
	database string = utils.GetEnvironment("database")
)

var (
	dsn    = fmt.Sprintf("%v:%v@/%v", username, password, database)
	dbcon  *sql.DB
	newErr error
)

func GetConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error : ", err)
	}

	return db, err
}
