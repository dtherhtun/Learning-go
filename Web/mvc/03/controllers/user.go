package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dtherhtun/Learning-go/Web/mvc/03/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	session *mongo.Collection
}

func NewUserController(s *mongo.Collection) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	filter := bson.M{"_id": params["id"]}

	u := models.User{}

	if err := uc.session.FindOne(context.TODO(), filter).Decode(&u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(u)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	result, err := uc.session.InsertOne(context.TODO(), u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(result)
		w.Write(response)
	}
}

func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	filter := bson.M{"_id": mux.Vars(r)["id"]}
	update := bson.M{"$set": &u}

	result, err := uc.session.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(result)
	w.Write(response)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	filter := bson.M{"_id": mux.Vars(r)["id"]}

	result, err := uc.session.DeleteOne(context.TODO(), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(result)
		w.Write(response)
	}
}
