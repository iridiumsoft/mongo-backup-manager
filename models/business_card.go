package models

import "gopkg.in/mgo.v2/bson"

type BusinessCard struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name,omitempty"`
	UserName    string        `json:"user_name,omitempty"`
	Designation string        `json:"designation,omitempty"`
	Address     string        `json:"address,omitempty"`
	Zip         string        `json:"zip,omitempty"`
	City        string        `json:"city,omitempty"`
	State       string        `json:"state,omitempty"`
	Email       string        `json:"email,omitempty"`
	User        string        `json:"user,omitempty"`
}
