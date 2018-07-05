package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Business struct {
	Id           bson.ObjectId       `json:"_id" bson:"_id,omitempty"`
	User         string              `json:"user,omitempty"`
	Client       bson.ObjectId       `json:"client" bson:"client,omitempty"` // TODO:: make sure the names are correct in both metas
	Provider     bson.ObjectId       `json:"provider" bson:"client,omitempty"`
	Product      bson.ObjectId       `json:"product" bson:"client,omitempty"`
	Split        interface{}         `json:"split,omitempty"`
	SplitUser    string              `json:"split_user,omitempty"`
	SplitValue   float64             `json:"split_value,omitempty"`
	PolicyNumber string              `json:"policy_number,omitempty"`
	Amount       float64             `json:"amount,omitempty"`
	Status       interface{}         `json:"status,omitempty"`
	Paid         int                 `json:"paid,omitempty"`
	Meta         interface{}         `json:"meta,omitempty"`
	V            BusinessVirtualData `json:"v,omitempty"`
	CreatedOn    time.Time           `json:"created_on,omitempty"`
	SubmitDate   time.Time           `json:"submit_date,omitempty"`
}

type BusinessVirtualData struct {
	User     string `json:"user,omitempty"` // Assocate First Name
	Client   string `json:"c,omitempty"`    // Client First Name
	Provider string `json:"prov,omitempty"` // Provider
	Product  string `json:"prod,omitempty"` // Product
}
