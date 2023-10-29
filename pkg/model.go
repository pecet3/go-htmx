package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDb() {
	db, err := sql.Open("postgres", "host=jakubpacewi.cz, port=5432, password=reksio1, dbname=db1")
	if err != nil {
		log.Fatal("error during connection to database")

	}
	fmt.Println(db)
	err = db.Ping()
	if err != nil {
		log.Fatal("error during ping database")

	}

}
