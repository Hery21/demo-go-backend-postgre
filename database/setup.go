package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	HOST    = ""
	PORT    = 
	DB_USER = ""
	DB_PASS = ""
	DB_NAME = ""
)

func InitDB() *sql.DB {
	// combine string
	// sslmode is whether we use ssl or not
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, DB_USER, DB_PASS, DB_NAME)

	// open connection to DB
	DB, err := sql.Open("postgres", dbInfo)

	// check if error
	if err != nil {
		fmt.Println("Error connecting to database")
		fmt.Println(err.Error())
	}

	return DB
}
