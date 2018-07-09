package controllers

import (
	"os"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

func (c *Controllers) InitBot() {
	jobs := cron.New()
	AppEnv := os.Getenv("APP_ENV")
	if AppEnv == "production" {
		logrus.Println("We are in Production Mode")
		jobs.AddFunc("@daily", func() {
			logrus.Println("Bot started on @Daily Basis")
			go c.DBDump(false)
		})
	}
}
