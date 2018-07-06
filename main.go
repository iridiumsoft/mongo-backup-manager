package main

import (
	"time"

	"github.com/iridiumsoft/mongo-backup-manager/controllers"
	"github.com/iridiumsoft/mongo-backup-manager/config"
	"github.com/iridiumsoft/mongo-backup-manager/app"

	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

func main() {

	// Custom Configuration
	configure()

	// Load Configuration file
	conf := config.FromFile("./config/config.json")

	// Initiate Database Connection
	db, err := initDB(conf)

	// If connection failed, keep trying after every other 10 seconds
	for err != nil {
		logrus.WithError(err).Error("Failed init connection to database. Trying again after 10 seconds")
		time.Sleep(time.Second * 10)
		db, err = initDB(conf)
	}

	// Create new Application
	application := app.New(db, conf)
	// Pass app and config to Controllers
	c := controllers.New(conf, application)

	// Start Bot
	c.InitBot();

	// Start Routing
	err = c.Routing()
	if err != nil {
		logrus.Error(err.Error())
	}
}

func configure() {
	// Set logger format
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func initDB(conf config.Main) (*mgo.Database, error) {
	session, err := mgo.Dial("mongodb://" + conf.DB.Host)
	db := session.DB(conf.DB.Name)
	return db, err
}
