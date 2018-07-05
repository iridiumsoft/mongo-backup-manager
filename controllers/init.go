package controllers

import (
	"time"

	"github.com/iridiumsoft/mongo-backup-manager/app"
	"github.com/iridiumsoft/mongo-backup-manager/config"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
)

// Controller implemented router of project
type Controllers struct {
	Config config.Main
	App    *app.App
	Gin    *gin.Engine
}

func initGin() *gin.Engine {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders: "Origin, Content-Type, Access-Control-Allow-Origin",
		ExposedHeaders: "",
		MaxAge:         50 * time.Second, // TODO:: Replace with appropriate Number
	}))
	return g
}

func New(config config.Main, application *app.App) *Controllers {
	return &Controllers{
		Config: config,
		Gin:    initGin(), // New Gin Context
		App:    application,
	}
}
