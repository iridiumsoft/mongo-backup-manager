package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Memos struct {
	Id                     bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User                   string        `json:"user,omitempty"`
	MemoId                 string        `json:"memo_id,omitempty"`
	Type                   int           `json:"type,omitempty"`
	SubType                int           `json:"subType,omitempty"`
	CreatedOn              time.Time     `json:"created_on,omitempty"`
	StartDate              time.Time     `json:"start_date,omitempty"`
	ExpiryDate             time.Time     `json:"expiry_date,omitempty"`
	UpdatedOn              time.Time     `json:"updated_on,omitempty"`
	Subject                string        `json:"subject,omitempty"`
	Message                string        `json:"message,omitempty"`
	PageId                 string        `json:"page_id,omitempty"`
	State                  []interface{} `json:"state,omitempty"`
	LicenseTypes           []interface{} `json:"license_types,omitempty"`
	AdditionalNotification []interface{} `json:"additional_notification,omitempty"`
	Users                  []interface{} `json:"users,omitempty"`
	Status                 int           `json:"status,omitempty"`
	Display                int           `json:"display,omitempty"`
	Rmd                    bool          `json:"rmd,omitempty"`
	Notified               bool          `json:"notified,omitempty"`
}
