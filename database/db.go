package database

import (
	"database/sql"
	"fmt"
	"mahasiswa/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetDB() (*sql.DB, error) {
	configuration := config.GetConfig()

	connect_string := configuration.DB_USERNAME + ":" + configuration.DB_PASSWORD + "@tcp(" + configuration.DB_HOST + ":" + configuration.DB_PORT + ")/" + configuration.DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", connect_string)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(err.Error())
	}

	return db, nil
}
