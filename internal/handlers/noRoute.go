package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoRouteHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{})
}
