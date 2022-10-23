package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	myRouter := http.NewServeMux()
	myRouter.HandleFunc("/", app.home)

	return myRouter
}
