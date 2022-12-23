package main

import (
	"fmt"
	"net/http"

	"github.com/HMadhav/CRM/handlers"
)

func main() {

	router := handlers.Routes()

	//Server configuration
	fmt.Println("Server is running on port: 8000")
	http.ListenAndServe(":8000", router)
}
