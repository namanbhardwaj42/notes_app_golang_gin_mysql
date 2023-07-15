package controllers

import (
	"fmt"
	"goApp/helpers"
	"goApp/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SignUp
func CreateUser(ctx *gin.Context) {

	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	if name == "" || email == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "One or more params are missing/wrong",
		})
	} else {
		err := models.DB_CreateUser(name, email, password)
		if err != nil {
			fmt.Println("Error : Create User!")
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Successfully Created",
			})
		}
	}

}

// Login

func Login(ctx *gin.Context) {

	currentUser := helpers.GetUserFromRequest(ctx)
	if currentUser == nil || currentUser.ID == 0 {

		email := ctx.PostForm("email")
		password := ctx.PostForm("password")

		user := models.UserCheck(email)

		if user != nil && user.Password == password {
			// Set Sesssion
			helpers.SessionSet(ctx, user.ID)

			session := sessions.Default(ctx)
			session_id := session.Get("id")

			ctx.JSON(http.StatusOK, gin.H{
				"sid": session_id,
			})

		} else {
			err := "Invalid Credentials!"
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
		}

	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "Already Logged In",
			"user_id": currentUser.ID,
		})
		return
	}

}

func Logout(ctx *gin.Context) {
	// Clear the session

	helpers.SessionClear(ctx)
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"status": "logged out",
	})
}
