package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/matcornic/hermes/v2"
)



func main(){
	userName:="Handu"
	userID:=122323
	VerificationCode:="dggsgfsgs"
	mode:="abc"
	fileId:="sdgsgdsgsd"
	WelcomeTempGenerator(userName, userID) 
	VerificationTempGenerator(userName , userID, VerificationCode) 
	RestoreIDTempGenerator(userName, userID) 
	FileDownloadTempGenerator(userName, mode, userID, fileId )

}