package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type CommissionType struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name      string        `json:"name,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
