package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Html struct {
	Id         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title      string        `json:"name,omitempty"`
	Slug       string        `json:"slug,omitempty"`
	Domain     string        `json:"domain,omitempty"`
	Email      string        `json:"email,omitempty"`
	User       string        `json:"user,omitempty"`
	Phone      string        `json:"phone,omitempty"`
	Career     string        `json:"career,omitempty"`
	Duplicated bool          `json:"duplicated,omitempty"`
	Visits     int64         `json:"visits,omitempty"`
	Date       time.Time     `json:"date,omitempty"`
}
