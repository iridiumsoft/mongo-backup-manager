package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Token struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User      string        `json:"user,omitempty"`
	Type      string        `json:"type,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
