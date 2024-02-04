package handlers

import (
	"net/http"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error Parsing the form", http.StatusInternalServerError)
			return
		}
		action := r.FormValue("videoAction")
		switch action {
		case "Convert":
			http.Redirect(w, r, "convert/index.html", http.StatusSeeOther)
		case "Cut":
			http.Redirect(w, r, "../cut/index.html", http.StatusSeeOther)
		case "Watermark":
			http.Redirect(w, r, "../watermark/index.html", http.StatusSeeOther)
		case "GIF":
			http.Redirect(w, r, "../gif/index.html", http.StatusSeeOther)
		default:
			http.Redirect(w, r, "/", http.StatusSeeOther)

		}
	}
} //Without caps it won't be exported

//so Basically it first gets and then posts
