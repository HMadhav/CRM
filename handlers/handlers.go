package handlers

import (
	"net/http"

	"github.com/HMadhav/CRM/controllers"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	// Register all the standard library debug endpoints.
	router.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", controllers.GetCustomer).Methods("GET")
	router.HandleFunc("/customers", controllers.AddCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", controllers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", controllers.DeleteCustomer).Methods("DELETE")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static"))))

	return router
}
