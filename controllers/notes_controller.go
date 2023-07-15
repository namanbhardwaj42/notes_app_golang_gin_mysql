package controllers

import (
	"fmt"
	"goApp/helpers"
	"goApp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NotesIndex(ctx *gin.Context) {

	currentUser := helpers.GetUserFromRequest(ctx)
	if currentUser == nil || currentUser.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "Unauthorized Access!",
		})
		return
	}
	notes := models.NotesAll(currentUser)
	ctx.JSON(http.StatusOK, gin.H{
		"notes": notes,
	})
}

func NotesCreate(ctx *gin.Context) {
	currentUser := helpers.GetUserFromRequest(ctx)
	if currentUser == nil || currentUser.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "Unauthorized Access!",
		})
		return
	}
	notes := ctx.PostForm("notes")

	models.NoteCreate(currentUser, notes)
	lastNoteID := models.LastNoteID(currentUser)
	ctx.JSON(http.StatusOK, gin.H{
		"id": lastNoteID.ID,
	})
}

func NotesShow(ctx *gin.Context) {

	currentUser := helpers.GetUserFromRequest(ctx)
	if currentUser == nil || currentUser.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "Unauthorized Access!",
		})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(currentUser, id)
	ctx.JSON(http.StatusOK, gin.H{
		"note": note,
	})
}

func NotesDelete(ctx *gin.Context) {
	currentUser := helpers.GetUserFromRequest(ctx)
	if currentUser == nil || currentUser.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "Unauthorized Access!",
		})
		return
	}
	idStr := ctx.PostForm("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	models.NotesMarkDelete(currentUser, id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})
}
