package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type AchArchive struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	User      string        `json:"user,omitempty"`
	Batch     bson.ObjectId `bson:"batch,omitempty"`
	Title     string        `json:"title"`
	AsCorp    string        `json:"ac_corp"`
	Type      string        `json:"type"`
	Account   string        `json:"account"`
	Routing   string        `json:"routing"`
	Ssn       string        `json:"ssn"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
