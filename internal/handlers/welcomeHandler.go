package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/middlewares"
	"net/http"
)

func WelcomeHandler(c *gin.Context) {

	if !middlewares.IsUserAuthenticated(c) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	} else {
		c.Redirect(http.StatusFound, "/dashboard")
	}

}
