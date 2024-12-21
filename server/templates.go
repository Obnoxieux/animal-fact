package server

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const RootLayout = "layout.gohtml"

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("templates", tmpl+".gohtml")
	basePath := filepath.Join("templates", RootLayout)

	templates, err := template.ParseFiles(basePath, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	err = templates.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
