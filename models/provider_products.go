package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type ProviderProduct struct {
	Id           bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User         string        `json:"user,omitempty"`
	Name         string        `json:"name,omitempty"`
	Provider     string        `json:"provider,omitempty"`
	Category     string        `json:"category,omitempty"`
	Profile      string        `json:"profile,omitempty"`
	Bonus        int           `json:"bonus,omitempty"`
	Gross        int           `json:"gross,omitempty"`
	Factor       int           `json:"factor,omitempty"`
	Logo         string        `json:"logo,omitempty"`
	License      int           `json:"license,omitempty"`
	Renewal      int           `json:"renewal,omitempty"`
	LicenseTypes bson.ObjectId `json:"license_types" bson:"_id,omitempty"`
	CreatedOn    time.Time     `json:"created_on,omitempty"`
}
