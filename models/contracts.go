package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Contract struct {
	Id       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User     string        `json:"user,omitempty"`
	Provider string        `json:"provider,omitempty"`
	File     string        `json:"file,omitempty"`
	Status   interface{}   `json:"status,omitempty"`
	Date     time.Time     `json:"date,omitempty"`
}
