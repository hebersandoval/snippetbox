package main

import (
	"log"
	"net/http"
)

// home is a handler function which writes a byte slice as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// showSnippet displays specific snippets.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func main() {
	// Initialize a new servemux, then register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)

	// Start a new web server and pass the TCP network address to listen on and the servemux just created.
	// If http.ListenAndServe() returns an error, an error will be thrown.
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
