package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Task struct {
	Id    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Value interface{}   `json:"value,omitempty"`
	Call  string        `json:"call,omitempty"`
	When  time.Time     `json:"when,omitempty"`
}
