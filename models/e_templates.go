package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ETemplate struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Type        string        `json:"type,omitempty"`
	Text        string        `json:"text,omitempty"`
	TextColor   string        `json:"text_color,omitempty"`
	TextHColor  string        `json:"text_h_color,omitempty"`
	BColor      string        `json:"b_color,omitempty"`
	BackC       string        `json:"back_c,omitempty"`
	BackColor   string        `json:"back_color,omitempty"`
	FotText     string        `json:"fot_text,omitempty"`
	TColor      string        `json:"t_color,omitempty"`
	IsText      bool          `json:"isText,omitempty"`
	ShowFooter  bool          `json:"showFooter,omitempty"`
	ShowHeader  bool          `json:"showHeader,omitempty"`
	User        string        `json:"user,omitempty"`
	Date        time.Time     `json:"date,omitempty"`
}
