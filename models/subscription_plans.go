package models

import (
	"gopkg.in/mgo.v2/bson"
)

type SubscriptionPlan struct {
	Id              bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title           string        `json:"title,omitempty"`
	Name            string        `json:"name,omitempty"`
	Price           float64       `json:"price,omitempty"`
	RecruitingValue int64         `json:"recruiting_value,omitempty"`
	Active          int           `json:"active,omitempty"`
}
