package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type LbPoint struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Points      int           `json:"points,omitempty"`
	Title       string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	User        string        `json:"user,omitempty"`
	CreatedOn   time.Time     `json:"created_on,omitempty"`
}
