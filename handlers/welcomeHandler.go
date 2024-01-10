package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func WelcomeHandler(templatePath, pageDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join(pageDir, templatePath)
		tmpl, err := template.ParseFiles(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
