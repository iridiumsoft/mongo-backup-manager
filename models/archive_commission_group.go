package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ArchiveCommissionGroup struct {
	Id      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Subject string        `json:"subject,omitempty"`
	Date    time.Time     `json:"date,omitempty"`
	Data    interface{}   `json:"data,omitempty"`
}
