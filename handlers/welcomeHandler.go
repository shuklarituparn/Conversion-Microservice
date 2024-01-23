package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WelcomeHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{})

}

//TODO: To make a good welcome page and then add the login link there
