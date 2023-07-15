package middlewares

import (
	"goApp/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get session
		session := sessions.Default(c)
		sessionID := session.Get("id")

		var user *models.UserS
		userPresent := true
		if sessionID == nil {
			userPresent = false
		} else {
			user = models.UserFind(sessionID.(uint64))
			userPresent = (user.ID > 0)
		}

		if userPresent {
			c.Set("user_id", user.ID)
			c.Set("email", user.Email)
		}
		c.Next()
	}
}
