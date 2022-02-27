package routes

import (
	"github.com/gorilla/mux"
	"majoo-be-test/service"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", services.Login).Methods("POST")
	router.HandleFunc("/report", services.Report).Methods("GET")
	return router
}
