package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type User struct {
	Id              bson.ObjectId       `json:"_id" bson:"_id,omitempty"`
	FirstName       string              `json:"first_name,omitempty"`
	LastName        string              `json:"last_name,omitempty"`
	DisplayName     string              `json:"display_name,omitempty"`
	SpouseFirstName string              `json:"spouse_first_name,omitempty"`
	SpouseLastName  string              `json:"spouse_last_name,omitempty"`
	Email           string              `json:"email,omitempty"`
	UserName        string              `json:"user_name,omitempty"`
	About           string              `json:"about,omitempty"`
	Address         string              `json:"address,omitempty"`
	Address2        string              `json:"address2,omitempty"`
	Altphone        string              `json:"altphone,omitempty"`
	Phone           string              `json:"phone,omitempty"`
	City            string              `json:"city,omitempty"`
	State           string              `json:"state,omitempty"`
	Zip             string              `json:"zip,omitempty"`
	Country         string              `json:"country,omitempty"`
	Code            string              `json:"code,omitempty"`
	Recruiter       string              `json:"recruiter,omitempty"`
	Dp              string              `json:"dp,omitempty"`
	ApiKey          string              `json:"api_key,omitempty"`
	Type            string              `json:"type,omitempty"`
	Role            string              `json:"role,omitempty"`
	Plan            string              `json:"plan,omitempty"`
	ComPlan         string              `json:"com_plan,omitempty"`
	Level           string              `json:"level,omitempty"`
	Token           string              `json:"token,omitempty"`
	Bank            string              `json:"bank,omitempty"`
	Status          string              `json:"status,omitempty"`
	Verify          string              `json:"verify,omitempty"`
	ConstantContact string              `json:"constantContact,omitempty"`
	Coupon          string              `json:"coupon,omitempty"`
	Leader          string              `json:"leader,omitempty"`
	Rmd             string              `json:"rmd,omitempty"`
	TheRmd          string              `json:"the_rmd,omitempty"`
	LastLoginIp     string              `json:"last_login_ip,omitempty"`
	RegistrationIp  string              `json:"registration_ip,omitempty"`
	Tc              string              `json:"tc,omitempty"`
	Ssn             string              `json:"ssn,omitempty"`
	Completed       int                 `json:"completed,omitempty"`
	ClickSend       int64               `json:"clickSend,omitempty"`
	InvalidDate     bool                `json:"invalid_date,omitempty"`
	EnableDn        bool                `json:"enable_dn,omitempty"`
	CreatedOn       time.Time           `json:"created_on,omitempty"`
	CoLeader        CoLeader            `json:"coLeader,omitempty"`
	Subscription    UserSubscription    `json:"subscription,omitempty"`
	Payment         RegistrationPayment `json:"payment,omitempty"`
	Stripe          interface{}         `json:"stripe,omitempty"`
	License         interface{}         `json:"license,omitempty"`
	Dob             interface{}         `json:"dob,omitempty"`
	Answers         interface{}         `json:"answers,omitempty"`
}

func (u *User) Update(db *mgo.Database) {
	if u.Id != "" {
		db.C("users").Update(bson.M{"_id": bson.ObjectId(u.Id)}, u)
	}
}

func (u *User) Delete(db *mgo.Database) {
	if u.Id != "" {
		db.C("users").Remove(bson.M{"_id": bson.ObjectId(u.Id)})
	}
}

func (u *User) IsAdmin() bool {
	return u.Type != "normal"
}

type UserSubscription struct {
	Title string  `json:"title,omitempty"`
	Price float64 `json:"price,omitempty"`
}

type RegistrationPayment struct {
	Method string  `json:"method,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

type CoLeader struct {
	User       string      `json:"user,omitempty"`
	Percentage interface{} `json:"percentage,omitempty"`
}
