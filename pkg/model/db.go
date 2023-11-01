package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDb() {
	db, err := sql.Open("postgres", "user=user1 host=localhost port=5432 password=haslo dbname=db1 sslmode=disable")
	if err != nil {
		log.Fatal("error during connection to database")

	}
	fmt.Println(db)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)

	}

}
