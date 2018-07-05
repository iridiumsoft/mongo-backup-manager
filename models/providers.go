package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Provider struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User        string        `json:"user,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	Profile     string        `json:"profile,omitempty"`
	Logo        string        `json:"logo,omitempty"`
	Url         string        `json:"url,omitempty"`
	Category    string        `json:"category,omitempty"`
	AdvanceRate interface{}   `json:"advance_rate,omitempty"`
	Status      interface{}   `json:"status,omitempty"`
	UpdatedOn   time.Time     `json:"updated_on,omitempty"`
	CreatedOn   time.Time     `json:"created_on,omitempty"`
}
