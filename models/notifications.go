package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Notification struct {
	Id      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Message string        `json:"m,omitempty"`
	T       string        `json:"t,omitempty"`
	L       string        `json:"l,omitempty"`
	Type    string        `json:"type,omitempty"`
	To      string        `json:"to,omitempty"`
	Status  interface{}   `json:"status,omitempty"`
	Date    time.Time     `json:"date,omitempty"`
}
