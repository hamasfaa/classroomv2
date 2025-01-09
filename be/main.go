package main

import (
	"be/config"
	"be/controller/usercontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/login", usercontroller.Login)

	log.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
