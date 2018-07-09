package controllers

import (
	"net/http"

	"github.com/iridiumsoft/mongo-backup-manager/models"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (c *Controllers) Routing() error {

	// Ping
	c.Gin.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})

	// Load Views
	c.Gin.LoadHTMLGlob("views/*")

	// Set Static files folder
	c.Gin.Static("/assets", "./assets")

	// Backups
	c.Gin.GET("/backup", c.ManualBackup)
	c.Gin.GET("/backup/export/:collection", c.ExportCollection)
	c.Gin.GET("/backup/import", c.ImportView)
	c.Gin.POST("/backup/import", c.ImportCollection)
	c.Gin.GET("/backup/restore/:id", c.RestoreDB)
	// Dashboard
	c.Gin.GET("/dashboard", func(context *gin.Context) {
		context.HTML(http.StatusOK, "dashboard.html", bson.M{})
	})
	// Make sure every requests is authenticated
	authorize := c.Gin.Group("/auth")
	authorize.Use(c.isLogin)
	authorize.GET("user", func(context *gin.Context) {
		context.String(http.StatusOK, "Ok")
	})
	return c.Gin.Run(":" + c.Config.Port)
}

func (c *Controllers) isLogin(ctx *gin.Context) {
	token := ctx.GetHeader("t")
	isLogin, User := c.getUserByToken(token)
	if !isLogin {
		ctx.JSON(http.StatusForbidden, gin.H{
			"loginfailed": true,
			"status":      false,
			"msg":         "Invalid or Empty Token",
		})
		return
	}
	ctx.Set("user", User)
	ctx.Next()
}

func (c *Controllers) isAdmin(ctx *gin.Context) {
	User, exists := ctx.Get("user")
	if exists {
		user := User.(models.User)
		if user.Type == "normal" {
			ctx.Abort()
			return
		}
		ctx.Next()
		return
	}
	ctx.Abort()
}

func (c *Controllers) getUserByToken(token string) (isLogin bool, User models.User) {
	var tokenObject models.Token;
	c.App.DB.C("tokens").FindId(token).One(&tokenObject)
	if tokenObject.Id != "" {
		c.App.DB.C("users").Find(bson.M{"user_name": tokenObject.User}).One(&User)
		if User.UserName != "" {
			isLogin = true
		}
	}
	return isLogin, User
}
