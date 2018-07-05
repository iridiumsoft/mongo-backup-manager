package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type License struct {
	Id         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User       string        `json:"user,omitempty"`
	Number     int           `json:"number,omitempty"`
	Expiration string        `json:"expiration,omitempty"`
	Types      interface{}   `json:"types,omitempty"`
	Status     int           `json:"status,omitempty"`
	CreatedOn  time.Time     `json:"created_on,omitempty"`
}
