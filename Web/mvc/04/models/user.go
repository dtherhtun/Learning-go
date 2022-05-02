package models

import "time"

type User struct {
	//Id       string `json:"id" bson:"_id"`
	UserName string `json:"username" bson:"username"`
	Password []byte `json:"password" bson:"password"`
	Name     string `json:"name" bson:"name"`
	Role     string `json:"role" bson:"role"`
	Gender   string `json:"gender" bson:"gender"`
	Age      string `json:"age" bson:"age"`
}

type Session struct {
	UserName     string
	LastActivity time.Time
}
