package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/ID"
	"github.com/shuklarituparn/Conversion-Microservice/configs"
	"github.com/shuklarituparn/Conversion-Microservice/user_sessions"
	"log"
	_ "log"
	"net/http"
)

var state = ID.ReturnID()

func Login(c *gin.Context) {

	if !userAlreadyLoggedIn(c) {
		conf := configs.Config()
		url := conf.AuthCodeURL(state)

		c.Redirect(http.StatusFound, url)
	} else {
		c.Redirect(http.StatusFound, "/dashboard")
	}

}

func userAlreadyLoggedIn(c *gin.Context) bool {
	session, err := user_sessions.Store.Get(c.Request, "Logged_Session")
	if err != nil {

		log.Println("Error finding session")
	}
	auth, ok := session.Values["authenticated"].(bool) //type assertion for the value from the session
	return ok && auth
}
