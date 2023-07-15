/*

						:::: Notes App ::::
	Tech & Tools : Golang , Gin , Gorm, Mysql , Xampp Server(for MySql)
	Server runs on port `:3000`
	DB Connection and settings --> models -> Setup

	Script Written by Naman Bhardwaj
	Contact - naman.bhardwaj42@gmail.com
*/

package main

import (
	"goApp/controllers"
	"goApp/middlewares"
	"goApp/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.New()

	router.SetTrustedProxies(nil)

	models.ConnectDB()
	models.DB_Table_Migrate()

	//Session Init
	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("notes", store))

	router.Use(middlewares.AuthenticateUser())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/signup", controllers.CreateUser)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)

	router.POST("/notes", controllers.NotesCreate)
	router.GET("/notes", controllers.NotesIndex)

	router.DELETE("/notes", controllers.NotesDelete)

	router.Run(":3000")
}
