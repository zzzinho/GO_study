package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/study/api_structure/config"
)

var DB *sql.DB

func Connect() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.DB_USER, config.DB_PASSWORD, config.DB_NAME)

	db, _ := sql.Open("postgres", dbinfo)
	err := db.Ping()
	if err != nil {
		log.Fatal("Error: Could not make a connection with the database")
	}
	DB = db

	CreateUserTable()
}
