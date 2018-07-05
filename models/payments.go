package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Payments struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User      string        `json:"user,omitempty"`
	Paid      string        `json:"paid,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
