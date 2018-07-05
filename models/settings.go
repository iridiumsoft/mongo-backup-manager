package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Setting struct {
	Id    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Value bson.M        `json:"n,omitempty"`
	Name  string        `json:"v,omitempty"`
}
