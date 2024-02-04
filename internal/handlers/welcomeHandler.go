package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WelcomeHandler(c *gin.Context) {

	if !IsUserAuthenticated(c) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	} else {
		c.Redirect(http.StatusFound, "/dashboard")
	}

}
