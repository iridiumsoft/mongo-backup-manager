package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ProductProviderCategory struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name      string        `json:"name,omitempty"`
	User      string        `json:"user,omitempty"`
	F         string        `json:"f,omitempty"`
	Bv        int           `json:"bv,omitempty"`
	Year1     int           `json:"year1,omitempty"`
	Year2     int           `json:"year2,omitempty"`
	Year3     int           `json:"year3,omitempty"`
	Year1_    interface{}   `json:"_year1,omitempty"`
	Year2_    interface{}   `json:"_year2,omitempty"`
	Year3_    interface{}   `json:"_year3,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
}
