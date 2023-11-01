package main

import (
	"fmt"

	model "github.com/pecet3/go+htmx/pkg/model"
)

func main() {
	model.ConnectDb()

	// model.CreateTask("test")

	tasks, err := model.GetAllTasks()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tasks)

}
