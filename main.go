package main

import (
	"mahasiswa/database"
	"mahasiswa/router"
)

func main() {
	db, err := database.GetDB()

	if err != nil {
		panic(err.Error())
	}
	
	e := router.Init(db)

	e.Run(":8080")
}
