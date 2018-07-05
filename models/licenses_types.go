package models

import (
	"gopkg.in/mgo.v2/bson"
)

type LicensesType struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User        string        `json:"user,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
}
