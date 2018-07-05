package app

import (
	"github.com/iridiumsoft/mongo-backup-manager/util"
	"github.com/iridiumsoft/mongo-backup-manager/config"

	"gopkg.in/mgo.v2"
)

// App struct implemented domain layer.

type App struct {
	DB        *mgo.Database
	DbService *util.DbService
	S3Service *util.S3Service
}

func New(db *mgo.Database, config config.Main) *App {
	r := &App{
		DB:        db,
		DbService: util.DbService{}.New(db),
		S3Service: util.S3Service{}.New(config),
	}
	return r
}
