package handlers

import (
	"fmt"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "YES IT HITS")
}
