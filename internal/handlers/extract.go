package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"log"
	"net/http"
)

func Screenshot(c *gin.Context) {

	session, err := user_sessions.Store.Get(c.Request, "Logged_Session") //getting the session from the session store
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	userName, ok := session.Values["userName"].(string)
	if !ok {
		log.Println("Error finding userID from the sessions")
	}
	userPicture, ok := session.Values["userPhoto"].(string)
	if !ok {
		log.Println("Error finding userID from the sessions")
	}

	userId, ok := session.Values["UserId"].(int)
	if !ok {
		log.Println("Error resolving the userId from the sessions")
	}
	db := user_database.ReturnDbInstance()
	VerfiedCheck, _ := user_database.IsVerified(db, userId)
	if !VerfiedCheck {
		c.HTML(http.StatusFound, "email_redirect.html", gin.H{})

	} else {
		c.HTML(http.StatusOK, "screenshot.html", gin.H{
			"userName":    userName,
			"userpicture": userPicture,
		})
	}
}
