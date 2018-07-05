package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Ama struct {
	Id   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User string        `json:"user,omitempty"`
	Ip   string        `json:"ip,omitempty"`
	Name string        `json:"name"`
	Date time.Time     `json:"date,omitempty"`
}
