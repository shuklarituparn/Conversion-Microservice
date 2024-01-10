package handlers

import (
	"fmt"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintf(w, "You made a post request to the upload Endpoint")
	}
	if r.Method == "GET" {
		fmt.Fprintf(w, "You made a get request to the upload Endpoint")
	}
	fmt.Println(r)
}

//So the post request is getting made to the endpoint
