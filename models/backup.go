package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Backup struct {
	Id   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name string        `json:"name"`
	Date time.Time     `json:"date"`
}
