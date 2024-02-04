package handlers

import (
	"fmt"
	"github.com/shuklarituparn/Conversion-Microservice/ID"
	"io"
	"net/http"
	"os"
	"strings"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(20 << 20) //Max Memory is 20Mb

	file, _, err := r.FormFile("Myfile")
	if err != nil {
		fmt.Fprintf(w, "Error reading your file")
		return
	}
	defer file.Close()
	//Newfile, _ := os.CreateTemp("userfiles", ".mp4")
	errChangingDir := os.Chdir("userfiles")
	if errChangingDir != nil {
		fmt.Fprintf(w, "Error creating the dir")
	}
	userchoice := r.FormValue("userchoice")
	Newfile, _ := os.Create("file1" + ID.ReturnID() + ".mp4")
	//Newfile, _ := os.Create("file1." + userchoice) //To convert it with the userchoice. I guess I'll need to save it

	fileBytes, _ := io.ReadAll(file)

	mimeType := http.DetectContentType(fileBytes)

	fmt.Fprintf(w, "Userchoice is %s", userchoice)

	if strings.HasPrefix(mimeType, "video/mp4") {
		Newfile.Write(fileBytes)
		fmt.Fprintf(w, "file successfully uploaded")
	} else {
		fmt.Fprintf(w, "BRO I said video only")
	}
	if err != nil {

		fmt.Println("Error reading the btes from file")
	}

}

//TODO: Now need to create a producer for this message. Can make a folder for init kafka and call the functions here
//TODO: Also left to add the Mongo Upload and good UI

//TODO: Need to check buffalo and if it makes it any better than from scratch

//So the post request is getting made to the endpoint

//and I am getting the data

//Need to process the data here
