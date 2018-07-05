package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserHirarchy struct {
	Id         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name       string        `json:"name,omitempty"`
	List       []interface{} `json:"list,omitempty"`
	UpdateNeed bool          `json:"update_need,omitempty"`
	UpdatedOn  time.Time     `json:"updated_on,omitempty"`
	CreatedOn  time.Time     `json:"created_on,omitempty"`
}
