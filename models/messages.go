package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Message struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Group     bson.ObjectId `json:"group" bson:"group,omitempty"`
	Text      string        `json:"text,omitempty"`
	User      interface{}   `json:"user,omitempty"`
	Status    int           `json:"status,omitempty"`
	CreatedAt time.Time     `json:"createdAt,omitempty"`
}
