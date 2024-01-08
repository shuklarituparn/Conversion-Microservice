//package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//)
//
//// downloadFileFunc isn't used in this snippet, but you can define its logic
//// and add a route for it if needed.
//func downloadFileFunc(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Download File Function called")
//}
//
//// handlerFunc checks for the specific query parameter and responds accordingly.
//func handlerFunc(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Query().Get("videoaction") == "convert" {
//
//		fmt.Fprintf(w, "Conversion page")
//	} else {
//		// If the specific query is not present, you can redirect to the home page or handle differently
//		fmt.Fprintf(w, "Hello, you're at the home page")
//	}
//}
//
//func main() {
//
//	fs := http.FileServer(http.Dir("static"))
//	http.Handle("/static/", http.StripPrefix("/static/", fs))
//
//	http.HandleFunc("/", handlerFunc)
//
//	fmt.Println("Server is starting on port 8085...")
//	if err := http.ListenAndServe(":8085", nil); err != nil {
//		log.Fatal(err)
//	}
//}

//TO CHECK THE OS.EXEC COMMANDS

//package main
//
//import (
//	"fmt"
//	"github.com/google/uuid"
//)
//
//func main() {
//	u, err := uuid.NewRandom()
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(u.String()) //more clearer and better to use
//
//}


