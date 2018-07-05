package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type CommissionBatche struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Description string        `json:"description,omitempty"`
	Users       []interface{} `json:"users,omitempty"`
	User        string        `json:"user,omitempty"`
	EmailSent   bool          `json:"email_sent,omitempty"`
	Old         int           `json:"old,omitempty"`
	Status      int           `json:"status,omitempty"`
	Date        time.Time     `json:"date,omitempty"`
	CreatedOn   time.Time     `json:"created_on,omitempty"`
}
