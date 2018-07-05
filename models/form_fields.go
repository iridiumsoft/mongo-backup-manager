package models

import (
	"gopkg.in/mgo.v2/bson"
)

type FormField struct {
	Id     bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name   string        `json:"name,omitempty"`
	Title  string        `json:"title,omitempty"`
	Page   string        `json:"page,omitempty"`
	Type   string        `json:"type,omitempty"`
	Active int           `json:"active,omitempty"`
}
