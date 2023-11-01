package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {

}
func sendTask(w http.ResponseWriter) {

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
