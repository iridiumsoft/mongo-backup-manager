package controllers

import "github.com/robfig/cron"

func (c *Controllers) InitBot() {
	jobs := cron.New()

	jobs.AddFunc("@daily", func() {
		go c.DBDump(false)
	})

	go c.DBDump(false)
}
