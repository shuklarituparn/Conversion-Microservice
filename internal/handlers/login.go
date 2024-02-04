package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/ID"
	"github.com/shuklarituparn/Conversion-Microservice/configs"
	_ "log"
	"net/http"
)

var state = ID.ReturnID()

func Login(c *gin.Context) {

	conf := configs.Config()
	url := conf.AuthCodeURL(state)
	c.Redirect(http.StatusFound, url)

}
