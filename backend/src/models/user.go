package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type (
	User struct {
		ID          bsontype.Type `json:"id" bson:"_id,omitempty"`
		Email       string        `json:"email,omitempty" bson:"email"`
		Username    string        `json:"username" bson:"username"`
		Password    string        `json:"password" bson:"password"`
		DateCreated time.Time     `json:"datecreated" bson:"datecreated"`
		Token       string        `json:"tokens,omitempty" bson:"-"`
	}
)
