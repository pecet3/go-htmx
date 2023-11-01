package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func ConnectDb() error {
	var err error
	db, err = sql.Open("postgres", "user=user1 host=localhost port=5432 password=haslo dbname=db1 sslmode=disable")
	if err != nil {
		log.Fatal("error during connection to database")

	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
