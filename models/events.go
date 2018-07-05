package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Event struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Address     string        `json:"address,omitempty"`
	Thumbnail   string        `json:"thumbnail,omitempty"`
	LargeBanner string        `json:"large_banner,omitempty"`
	Url         string        `json:"url,omitempty"`
	User        string        `json:"user,omitempty"`
	Start       int64         `json:"start,omitempty"`
	End         int64         `json:"end,omitempty"`
	Price       float64       `json:"price,omitempty"`
	Points      int64         `json:"points,omitempty"`
	NewTab      bool          `json:"newTab,omitempty"`
	Time        time.Time     `json:"time,omitempty"`
}
