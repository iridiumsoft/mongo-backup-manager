package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Commission struct {
	Id         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Percentage int           `json:"percentage,omitempty"`
	Perc       int           `json:"perc,omitempty"`
	Advance    int           `json:"advance,omitempty"`
	Deducted   int           `json:"deducted,omitempty"`
	Business   interface{}   `json:"business,omitempty"`
	User       string        `json:"user,omitempty"`
	Type       string        `json:"type,omitempty"`
	CType      string        `json:"c_type,omitempty"`
	BaseShop   string        `json:"base_shop,omitempty"`
	Category   string        `json:"category,omitempty"`
	Batch      bson.ObjectId `json:"batch" bson:"_id,omitempty"`
	Status     interface{}   `json:"status,omitempty"`
	CreatedOn  time.Time     `json:"created_on,omitempty"`
}
