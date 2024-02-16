package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FileHistory(c *gin.Context) {
	c.HTML(http.StatusFound, "file_history.html", gin.H{})
}
