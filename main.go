package main

import (
	"fmt"
	"github.com/shuklarituparn/Conversion-Microservice/handlers"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/videoAction", handlers.RedirectHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)

	fmt.Println("Server is starting on port 8085...")
	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatal(err)
	}
}

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

//package main
//
//import (
//	"html/template"
//	"net/http"
//)
//
//var templates = template.Must(template.ParseGlob("templates/*.html"))
//
//func main() {
//	http.HandleFunc("/", )
//	http.HandleFunc("/login", loginHandler)
//
//	fs := http.FileServer(http.Dir("static"))
//	http.Handle("/static/", http.StripPrefix("/static/", fs))
//
//	http.ListenAndServe(":8080", nil)
//}
//
//func homeHandler(w http.ResponseWriter, r *http.Request) {
//	if !isLoggedIn(r) {
//		http.Redirect(w, r, "/login", http.StatusSeeOther)
//		return
//	}
//
//	// Serve home page for logged-in users
//}
//
//func loginHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "GET":
//		renderTemplate(w, "index.html", nil)
//	case "POST":
//		// Process login form
//		// Authenticate user
//		// Redirect to home page if successful
//	}
//}
//
//func isLoggedIn(r *http.Request) bool {
//	// Implement logic to check if the user is logged in
//	return false
//}
//
//func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
//	err := templates.ExecuteTemplate(w, tmpl, data)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}

//localhost:8085/?videoaction=convert gives URL query  map[videoAction:[Convert]]
