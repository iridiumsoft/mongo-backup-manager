package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type FormVariable struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name      string        `json:"name,omitempty"`
	Label     string        `json:"label,omitempty"`
	P         string        `json:"p,omitempty"`
	Type      string        `json:"type,omitempty"`
	TType     string        `json:"_type,omitempty"`
	Options   interface{}   `json:"options,omitempty"`
	User      string        `json:"user,omitempty"`
	PlaceHold string        `json:"placehold,omitempty"`
	Active    int           `json:"active,omitempty"`
	CreatedOn time.Time     `json:"created_on,omitempty"`
	UpdatedOn time.Time     `json:"updated_on,omitempty"`
}
