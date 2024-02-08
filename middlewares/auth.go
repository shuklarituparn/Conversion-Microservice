package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !IsUserAuthenticated(c) {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}

func IsUserAuthenticated(c *gin.Context) bool {
	session, err := user_sessions.Store.Get(c.Request, "Logged_Session")
	if err != nil {

		return false
	}
	auth, ok := session.Values["authenticated"].(bool) //type assertion for the value from the session
	return ok && auth
}
