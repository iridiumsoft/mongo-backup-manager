package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ChargeBack struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User      string        `json:"user,omitempty"`
	Paid      interface{}   `json:"paid,omitempty"`
	Amount    int           `json:"amount,omitempty"`
	Business  bson.ObjectId `json:"business" bson:"_id,omitempty"`
	Reason    string        `json:"reason,omitempty"`
	Adjusted  bool          `json:"adjusted,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
