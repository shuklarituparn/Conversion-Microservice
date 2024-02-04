package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/configs"
	"github.com/shuklarituparn/Conversion-Microservice/internal/ID"
	_ "log"
	"net/http"
)

var state = ID.ReturnID()

func Login(c *gin.Context) {

	conf := configs.Config()
	url := conf.AuthCodeURL(state)
	c.Redirect(http.StatusFound, url)

}
