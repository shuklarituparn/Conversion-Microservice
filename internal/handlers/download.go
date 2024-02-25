package handlers

import "github.com/gin-gonic/gin"

func Download(c *gin.Context) {
	useridQ := c.Query("userid") //can match from the cookie if it is same, if not not allow
	modeFromQ := c.Query("mode") //the mode to finally decide which file which to get
	fileId := c.Query("fileid")  //the fileId to pull the file from the MongoDB database

	if useridQ == "" && modeFromQ == "" && fileId == "" {

	}

}
