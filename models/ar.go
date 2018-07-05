package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type AutoResponder struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Subject   string        `json:"subject,omitempty"`
	SendAfter string        `json:"send_after,omitempty"`
	Content   string        `json:"content,omitempty"`
	List      string        `json:"list,omitempty"`
	User      string        `json:"user,omitempty"`
	ReEmail   string        `json:"re_email,omitempty"`
	Template  string        `json:"template,omitempty"`
	Date      time.Time     `json:"date,omitempty"`
	Sent      int           `json:"sent,omitempty"`
}
