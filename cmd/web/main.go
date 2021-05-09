package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize a new servemux, then register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create a file server which serves files out of the "./ui/static" directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register the file server as the handler for all URL paths that start with "/static/" and strip the prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Start a new web server and pass the TCP network address to listen on and the servemux just created.
	// If http.ListenAndServe() returns an error, an error will be thrown.
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}


// Note: To run server from Windows OS -> go run ./cmd/web/.