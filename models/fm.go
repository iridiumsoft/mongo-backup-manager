package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type FileManager struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User      string        `json:"user,omitempty"`
	Location  string        `json:"location,omitempty"`
	Name      string        `json:"name,omitempty"`
	Mime      string        `json:"mime,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
