package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type EventRegistration struct {
	Id       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Event    interface{}   `json:"event,omitempty"`
	Status   int           `json:"status,omitempty"`
	StripeId string        `json:"stripe_id,omitempty"`
	Location string        `json:"location,omitempty"`
	Guests   interface{}   `json:"guests,omitempty"`
	User     interface{}   `json:"user,omitempty"`
	Date     time.Time     `json:"date,omitempty"`
}
