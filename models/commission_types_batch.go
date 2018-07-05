package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type CommissionTypesBatch struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User      string        `json:"name,omitempty"`
	Amount    int           `json:"amount,omitempty"`
	Business  interface{}   `json:"business,omitempty"`
	Type      interface{}   `json:"type,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
