package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Groups struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	User        string        `json:"user,omitempty"`
	Total       int64         `json:"total,omitempty"`
	CreatedOn   time.Time     `json:"created_on,omitempty"`
}
