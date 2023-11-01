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

func CreateTask(content string) error {
	statement := "INSERT INTO TASK (task, done) VALUES ($1, $2);"

	_, err := db.Query(statement, content, false)

	if err != nil {
		log.Fatal("error create task: ", err)
		return err
	}

	fmt.Println("created new record")

	return nil
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

func GetTaskById(Id uint64) (Task, error) {
	var task Task

	statement := "SELECT * FROM task WHERE id=$1"

	row, err := db.Query(statement, Id)

	if err != nil {
		log.Fatal("error during get taskby id: ", err)
		return task, err
	}

	for row.Next() {

		err := row.Scan(&task.Id, &task.Task, &task.Done)
		if err != nil {
			log.Fatal(err)
		}

	}
	return task, nil
}

func MarkDone(Id uint64) error {
	statement := "UPDATE task SET done=$2 WHERE id=$1;"

	task, err := GetTaskById(Id)
	if err != nil {
		log.Fatal("error during mark done: ", err)

	}

	_, err = db.Query(statement, Id, !task.Done)
	if err != nil {
		log.Fatal("error during mark done: ", err)
	}

	return nil
}

func DeleteById(Id uint64) error {
	statement := `DELETE FROM task WHERE id=$1;`

	_, err := db.Exec(statement, Id)
	return err
}
