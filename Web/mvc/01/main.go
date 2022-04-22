package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dtherhtun/Learning-go/Web/mvc/01/models"
	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    28,
		Id:     params["id"],
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = "007"

	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
