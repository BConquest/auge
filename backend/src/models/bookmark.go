package models

import (
	"time"
)

type (
	Bookmark struct {
		ID          string    `json:"id" bson:"_id,omitempty"`
		DateCreated time.Time `json:"datecreated" bson:"datecreated"`
		Link        string    `json:"link" bson:"link"`
		Tags        []string  `json:"tags" bson:"tags"`
		User        string    `json:"user" bson:"user"`
	}
)
