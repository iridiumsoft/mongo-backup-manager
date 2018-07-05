package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Lead struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	FirstName string        `json:"first_name,omitempty"`
	LastName  string        `json:"last_name,omitempty"`
	Email     string        `json:"email,omitempty"`
	Phone     string        `json:"phone,omitempty"`
	Number    string        `json:"number,omitempty"`
	Message   string        `json:"message,omitempty"`
	S         string        `json:"s,omitempty"`
	User      string        `json:"user,omitempty"`
	Date      time.Time     `json:"date,omitempty"`
	Other     interface{}   `json:"other,omitempty"`
}
