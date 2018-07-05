package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type SignedMemo struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User      string        `json:"user,omitempty"`
	Memo      bson.ObjectId `json:"memo" bson:"_id,omitempty"`
	Accept    int           `json:"accept,omitempty"`
	Date      time.Time     `json:"date,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
