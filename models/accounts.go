package models

import "gopkg.in/mgo.v2/bson"

type Account struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	User        string        `json:"user,omitempty"`
	Available   float64       `json:"available,omitempty"`
	ChargeBack  float64       `json:"chargeBack,omitempty"`
	LastPayment float64       `json:"last_payment,omitempty"`
}
