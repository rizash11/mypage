package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	tmpl, err := template.ParseFiles("./ui/html/home.page.html")
	if err != nil {
		return
	}

	tmpl.Execute(w, nil)
}
