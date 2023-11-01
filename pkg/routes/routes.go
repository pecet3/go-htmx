package routes

import (
	"html/template"
	"log"
	"net/http"

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

	err = tmp.ExecuteTemplate(w, "todos", todos)
	if err != nil {
		log.Fatal("error executing index.html: ", err)
	}

}

func markTaskDone(w http.ResponseWriter, r *http.Request) {

}
func deleteTask(w http.ResponseWriter, r *http.Request) {

}
func createTask(w http.ResponseWriter, r *http.Request) {

}
func SetupAndRun() {
	mux := mux.NewRouter()
	port := ":5000"

	mux.HandleFunc("/", index)

	mux.HandleFunc("/task/{id}", markTaskDone).Methods("PUT")
	mux.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
	mux.HandleFunc("/task", createTask).Methods("POST")

	log.Fatal(http.ListenAndServe(port, mux))
}
