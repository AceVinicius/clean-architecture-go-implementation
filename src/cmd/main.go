package main

import (
	"clean_architecture/src/adapter/api"
	"clean_architecture/src/adapter/database"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatal(err)
	}

	repository := database.NewTransactionRepositoryDb(db)
	webserver := api.NewWebServer(repository)

	webserver.Serve()
}
