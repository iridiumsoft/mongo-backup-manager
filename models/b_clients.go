package models

import "gopkg.in/mgo.v2/bson"

type Client struct {
	Id              bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	FirstName       string        `json:"first_name,omitempty"`
	LastName        string        `json:"last_name,omitempty"`
	Email           string        `json:"email,omitempty"`
	SpouseFirstName string        `json:"spouse_first_name,omitempty"`
	SpouseLastName  string        `json:"spouse_last_name,omitempty"`
	Address         string        `json:"address,omitempty"`
	Address2        string        `json:"address2,omitempty"`
	City            string        `json:"city,omitempty"`
	User            string        `json:"user,omitempty"`
}
