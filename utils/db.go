package utils

import (
	"database/sql"
	"fmt"
	"os"
)

// GetDBConnection - Opens a connection to the db and returns the a pointer to the db object
func GetDBConnection() *sql.DB {

	var (
		host     = "localhost"
		port     = 6543
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	return db
}
