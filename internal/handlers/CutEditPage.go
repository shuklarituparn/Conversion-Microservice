package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CutEditPage(c *gin.Context) {

	c.HTML(http.StatusFound, "cut_edit.html", gin.H{})

}
