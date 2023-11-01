package model

import (
	"fmt"
	"log"
)

type Task struct {
	Id   uint64 `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func CreateTask(content string) {
	statement := "INSERT INTO TASK (task, done) VALUES ($1, $2);"

	_, err := db.Query(statement, content, false)

	if err != nil {
		log.Fatal("error create task: ", err)
		return
	}

	fmt.Println("created new record")
}

func GetAllTasks() ([]Task, error) {
	tasks := []Task{}

	statement := "SELECT * FROM task;"

	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Task, &task.Done)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
