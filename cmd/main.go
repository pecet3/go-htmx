package main

import (
	"log"

	model "github.com/pecet3/go+htmx/pkg/model"
)

func main() {
	err := model.ConnectDb()

	if err != nil {
		log.Fatal(err)
	}
}
