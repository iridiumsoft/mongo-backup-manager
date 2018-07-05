package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Live struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Ip        string        `json:"ip,omitempty"`
	Socket    string        `json:"socket,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
