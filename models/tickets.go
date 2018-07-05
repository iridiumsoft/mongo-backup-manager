package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Ticket struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User      time.Time     `json:"user,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
