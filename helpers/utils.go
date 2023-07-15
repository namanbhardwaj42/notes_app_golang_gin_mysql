package helpers

import (
	"goApp/models"

	"github.com/gin-gonic/gin"
)

func GetUserFromRequest(ctx *gin.Context) *models.UserS {
	// Get user
	userID := ctx.GetUint64("user_id")
	var currentUser *models.UserS
	if userID > 0 {
		currentUser = models.UserFind(userID)
	} else {
		currentUser = nil
	}
	return currentUser
}
