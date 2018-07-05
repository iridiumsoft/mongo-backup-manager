package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Promotion struct {
	Id      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User    string        `json:"user,omitempty"`
	History []interface{} `json:"history,omitempty"`
}
