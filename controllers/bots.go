package controllers

import (
	"github.com/robfig/cron"
	"os"
	"github.com/sirupsen/logrus"
)

func (c *Controllers) InitBot() {
	jobs := cron.New()
	AppEnv := os.Getenv("APP_ENV")
	logrus.Print(AppEnv)

	jobs.AddFunc("@daily", func() {
		if AppEnv == "production" {
			go c.DBDump(false)
		}
	})
}
