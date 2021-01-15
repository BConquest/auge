package models

import (
	"time"
	//"go.mongodb.org/mongo-driver/bson/bsontype"
)

type (
	User struct {
		DateCreated time.Time `json:"datecreated" bson:"datecreated"`
		Email       string    `json:"email,omitempty" bson:"email"`
		ID          string    `json:"id" bson:"_id,omitempty"`
		Password    string    `json:"password" bson:"password"`
		Token       string    `json:"tokens,omitempty" bson:"-"`
		Type        string    `json:"accountType" bson:"accountType"`
		Username    string    `json:"username" bson:"username"`
	}
)
