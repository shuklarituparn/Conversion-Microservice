package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/user_sessions"
	"log"
	"net/http"
)

func Signout(c *gin.Context) {
	session, err := user_sessions.Store.Get(c.Request, "Logged_Session")
	if err != nil {
		log.Println("Error getting session:", err)
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	userName, ok := session.Values["userName"].(string)
	if !ok {
		log.Println("Error finding userID from the sessions")
	}
	userPicture, ok := session.Values["userPhoto"].(string)
	if !ok {
		log.Println("Error finding userID from the sessions")
	}

	session.Values["authenticated"] = false
	delete(session.Values, "userName")
	delete(session.Values, "userID")
	delete(session.Values, "userPhoto")
	session.Options.MaxAge = -1

	if err := session.Save(c.Request, c.Writer); err != nil {
		log.Println("Error saving session:", err)
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "signout.html", gin.H{
		"userImage": userPicture,
		"userName":  userName,
	})

}
