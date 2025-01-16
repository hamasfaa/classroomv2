package main

import (
	"be/config"
	"be/controller/dosencontroller"
	"be/controller/usercontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/login", usercontroller.Login)
	http.HandleFunc("/register", usercontroller.Register)
	http.HandleFunc("/logout", usercontroller.Logout)
	http.HandleFunc("/whoami", usercontroller.CurrentUser)

	http.HandleFunc("/dosen/", dosencontroller.Index)

	log.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
