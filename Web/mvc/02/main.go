package main

import (
	"net/http"

	"github.com/dtherhtun/Learning-go/Web/mvc/02/controllers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	uc := controllers.NewUserController()
	r.HandleFunc("/user/{id}", uc.GetUser).Methods("GET")
	r.HandleFunc("/user", uc.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", uc.Deleteuser).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
