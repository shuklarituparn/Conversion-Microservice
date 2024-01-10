package handlers

import (
	"fmt"
	"github.com/shuklarituparn/Conversion-Microservice/ID"
	"io"
	"log"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(100 << 20)
	fmt.Printf("HELLO FROM UPLOADS")
	filename := fmt.Sprintf("%s", ID.ReturnID())
	if r.Method == "POST" {
		fileFromUser, fileHeader, err := r.FormFile("Myfile")
		if err != nil {
			log.Fatalf("Error getting the file from USER due to %w", err)

		}
		defer fileFromUser.Close()
		fmt.Printf("Uploaded file %s", fileHeader.Filename)
		fmt.Printf("Size of the file %s", fileHeader.Size)
		userfile, err := os.Create(filename)
		defer userfile.Close()
		filebytes, _ := io.ReadAll(fileFromUser)
		userfile.Write(filebytes)
		fmt.Fprintf(w, "Successfully Uploaded File\n")

	}
}
