package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type PaidBonusPayment struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Business  interface{}   `json:"business,omitempty"`
	Amount    int           `json:"amount,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
