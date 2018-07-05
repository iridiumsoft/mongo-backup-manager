package models

import "gopkg.in/mgo.v2/bson"

type BusinessCategory struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	CategoryImg string        `json:"category_img,omitempty"`
	Parent      string        `json:"parent,omitempty"`
}
