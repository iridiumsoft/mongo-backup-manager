package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type BonusPayment struct {
	Id         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Business   bson.ObjectId `json:"business" bson:"_id,omitempty"`
	Type       bson.ObjectId `json:"type" bson:"_id,omitempty"`
	Amount     float64       `json:"amount,omitempty"`
	ChargeBack float64       `json:"chargeback,omitempty"`
	CreatedOn  time.Time     `json:"created_on,omitempty"`
}
