package routes

import (
	"github.com/gorilla/mux"
	"github.com/santhosh1188/employee-api/handlers"
)

func RegisterEmployeeRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/employees", handlers.GetEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", handlers.GetEmployee).Methods("GET")
	router.HandleFunc("/employees", handlers.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", handlers.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employees/{id}", handlers.DeleteEmployee).Methods("DELETE")
	return router
}
