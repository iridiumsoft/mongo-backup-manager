package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MsgGroup struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Users     []interface{} `json:"users,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
