package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"log"
	"net/http"
)

func Restore(c *gin.Context) {
	session, err := user_sessions.Store.Get(c.Request, "Logged_Session") //getting the session from the session store
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	userId, ok := session.Values["UserId"].(int) //this is needed to make sure, user doesn't delete someone else's account
	if !ok {
		log.Println("Error resolving the userId from the sessions")
	}
	db := user_database.ReturnDbInstance()
	var user models.User
	result := db.Unscoped().First(&user, userId)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
		log.Println("herere", result.Error)
	}
	secureKey := user.RestoreSecureKey
	c.HTML(http.StatusFound, "restore.html", gin.H{
		"secureCode": secureKey,
	})
}
