package main

import (
	"github.com/gorilla/mux"
	"go-jwt-mux/controllers/authcontroller"
	"go-jwt-mux/models"
	"log"
	"net/http"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
