package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Level struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	Pos         string        `json:"pos,omitempty"`
	Percentage  int64         `json:"percentage,omitempty"`
	Production  int64         `json:"production,omitempty"`
	Downlines   int64         `json:"downlines,omitempty"`
	User        string        `json:"user,omitempty"`
	ComPlan     string        `json:"com_plan,omitempty"` // TODO:: convert to ObjectId
	Date        time.Time     `json:"date,omitempty"`
}
