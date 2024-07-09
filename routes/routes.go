package routes

import (
	"github.com/DevMehta22/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/emps",controller.GetAllEmps).Methods("GET")
	router.HandleFunc("/api/emps/{id}",controller.GetEmpByID).Methods("GET")
	router.HandleFunc("/api/emps",controller.AddEmp).Methods("POST")
	router.HandleFunc("/api/emps/{id}",controller.UpdateEmp).Methods("PUT")
	router.HandleFunc("/api/emps/{id}",controller.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/emps",controller.DeleteAll).Methods("DELETE")
	
	return router

}