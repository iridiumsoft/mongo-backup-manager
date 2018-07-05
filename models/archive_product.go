package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ArchiveProduct struct {
	Id      bson.ObjectId      `json:"_id" bson:"_id,omitempty"`
	Subject string             `json:"subject,omitempty"`
	Date    time.Time          `json:"date,omitempty"`
	Data    ArchiveProductData `json:"data,omitempty"`
}

type ArchiveProductData struct {
	Factor float64 `json:"factor,omitempty"`
	Gross  float64 `json:"gross,omitempty"`
	Bonus  float64 `json:"bonus,omitempty"`
}
