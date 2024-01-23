package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{})
}
