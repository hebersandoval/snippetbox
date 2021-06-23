package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// Initialize a new servemux, then register the home function as the handler for the "/" URL pattern.
	// Swap the route declarations to use the application struct's methods as the handler functions.
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Create a file server which serves files out of the "./ui/static" directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// Register the file server as the handler for all URL paths that start with "/static/" and strip the prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}