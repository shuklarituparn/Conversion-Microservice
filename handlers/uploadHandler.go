package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintf(w, "You made a post request to the upload Endpoint")
	}
	if r.Method == "GET" {
		fmt.Fprintf(w, "You made a get request to the upload Endpoint")
	}
	p, _ := io.ReadAll(r.Body)
	fmt.Printf("%s\n", p)
}

//So the post request is getting made to the endpoint

//and I am getting the data
