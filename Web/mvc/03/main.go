package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dtherhtun/Learning-go/Web/mvc/03/controllers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := mux.NewRouter()
	uc := controllers.NewUserController(getSession())
	r.HandleFunc("/user/{id}", uc.GetUser).Methods("GET")
	r.HandleFunc("/user", uc.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", uc.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", uc.DeleteUser).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}

func getSession() *mongo.Collection {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb+srv://m001-student:m001-mongodb-basic@sandbox.4xpev.mongodb.net/go-web-dev-db").SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
	}
	collection := client.Database("go-web-dev-db").Collection("users")

	return collection
}
