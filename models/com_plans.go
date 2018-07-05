package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ComPlan struct {
	Id                   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name                 string        `json:"name,omitempty"`
	Description          string        `json:"description,omitempty"`
	GenerationPercentage int           `json:"generation_percentage,omitempty"`
	Default              int           `json:"default,omitempty"`
	Generations          []interface{} `json:"generations,omitempty"`
	CreatedOn            time.Time     `json:"created_on,omitempty"`
}
