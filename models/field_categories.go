package models

import (
	"gopkg.in/mgo.v2/bson"
)

type FieldCategory struct {
	Id     bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Status int           `json:"status,omitempty"`
	Title  string        `json:"title,omitempty"`
	User   string        `json:"user,omitempty"`
}
