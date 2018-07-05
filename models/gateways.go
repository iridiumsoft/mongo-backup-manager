package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Gateway struct {
	Id      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name    string        `json:"name,omitempty"`
	Server  string        `json:"server,omitempty"`
	Country string        `json:"country,omitempty"`
	User    string        `json:"user,omitempty"`
	Date    time.Time     `json:"date,omitempty"`
}
