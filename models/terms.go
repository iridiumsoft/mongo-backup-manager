package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Term struct {
	Id      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User    string        `json:"user,omitempty"`
	Content string        `json:"content,omitempty"`
}
