package models

import (
	"gopkg.in/mgo.v2/bson"
)

type UserMeta struct {
	Id   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User string        `json:"user,omitempty"`
	Bank interface{}   `json:"bank,omitempty"`
}
