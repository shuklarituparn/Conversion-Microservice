package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"log"
	"net/http"
)

func AccountDeleteConf(c *gin.Context) {
	session, err := user_sessions.Store.Get(c.Request, "Logged_Session") //getting the session from the session store
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	userId, ok := session.Values["UserId"].(int)
	if !ok {
		log.Println("Error resolving the userId from the sessions")
	}
	c.HTML(http.StatusFound, "delete_conf.html", gin.H{
		"userID": userId,
	})
}
