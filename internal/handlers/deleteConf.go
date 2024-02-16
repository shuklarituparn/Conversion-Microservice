package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccountDeleteConf(c *gin.Context) {
	c.HTML(http.StatusFound, "delete_conf.html", gin.H{})
}
