package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Contact struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	FirstName string        `json:"first_name,omitempty"`
	LastName  string        `json:"last_name,omitempty"`
	Email     string        `json:"email,omitempty"`
	Phone     string        `json:"phone,omitempty"`
	User      string        `json:"user,omitempty"`
	Group     string        `json:"group,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
