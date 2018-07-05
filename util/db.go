package util

import (
	"github.com/iridiumsoft/mongo-backup-manager/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
)

type DbService struct {
	DB *mgo.Database
}

func (d DbService) New(db *mgo.Database) *DbService {
	d.DB = db
	return &d
}

func (d *DbService) GetSetting(name string) bson.M {
	var setting models.Setting
	d.DB.C("settings").Find(models.Setting{Name: name}).Select("n").One(&setting)
	return setting.Value
}

func (d *DbService) UpdateSetting(name string, data bson.M) {
	d.DB.C("settings").Upsert(models.Setting{Name: name}, models.Setting{Name: name, Value: data})
}

func (d *DbService) IsObjectId(v interface{}) bool {
	t := reflect.TypeOf(v)
	// TODO:: Make sure 13 is the correct length
	return t != nil && t.Kind() == reflect.String && len(t.String()) == 13
}
