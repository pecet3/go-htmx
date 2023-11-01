package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pecet3/go+htmx/pkg/model"
)

func index(w http.ResponseWriter, r *http.Request) {
	todos, err := model.GetAllTasks()
	if err != nil {
		log.Fatal("error during get all tasks: ", err)
	}

	tmp := template.Must(template.ParseFiles("./pkg/templates/index.html"))

	err = tmp.Execute(w, todos)
	if err != nil {
		log.Fatal("error executing index.html: ", err)
	}

}
func sendTasks(w http.ResponseWriter) {
	todos, err := model.GetAllTasks()
	if err != nil {
		log.Fatal("error during get all tasks: ", err)
	}

	tmp := template.Must(template.ParseFiles("./pkg/templates/index.html"))

	err = tmp.ExecuteTemplate(w, "Todos", todos)
	if err != nil {
		log.Fatal("error executing index.html: ", err)
	}

}

func markTaskDone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(vars, 10, 64)
	fmt.Println(id)
	if err != nil {
		log.Fatal("error during parsing vars: ", err)
	}

	err = model.MarkDone(id)
	if err != nil {
		log.Fatal("error during mark done: ", err)
	}
	sendTasks(w)
}
func deleteTask(w http.ResponseWriter, r *http.Request) {

}
func createTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("error during parsing form data: ", err)
	}

	err = model.CreateTask(r.FormValue("task"))
	if err != nil {
		log.Fatal("error during creating task:", err)
	}

	sendTasks(w)
}
func SetupAndRun() {
	mux := mux.NewRouter()
	port := ":5000"

	mux.HandleFunc("/", index)

	mux.HandleFunc("/task/{id}", markTaskDone).Methods("PUT")
	mux.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
	mux.HandleFunc("/task/create", createTask).Methods("POST")

	log.Fatal(http.ListenAndServe(port, mux))
}
