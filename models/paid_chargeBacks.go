package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type PaidChargeBack struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Batch       bson.ObjectId `json:"batch" bson:"_id,omitempty"`
	ChargeBacks interface{}   `json:"chargebacks" bson:"_id,omitempty"`
	User        string        `json:"user,omitempty"`
	Amount      float64       `json:"amount,omitempty"`
	CreatedOn   time.Time     `json:"created_on,omitempty"`
	Date        time.Time     `json:"date,omitempty"`
}
