package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gorilla/sessions"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"log"
	"net/http"
)

func Dashboard(c *gin.Context) {

	session, err := user_sessions.Store.Get(c.Request, "Logged_Session") //getting the session from the session store
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	userPic, ok := session.Values["userPhoto"].(string)
	if !ok {
		log.Println("Error finding userPic from the sessions")
	}
	userName, ok := session.Values["userName"].(string)
	log.Println(userName)
	if !ok {
		log.Println("Error finding userName from the sessions")
	}
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"userpicture": userPic,
		"userName":    userName,
	})

}
