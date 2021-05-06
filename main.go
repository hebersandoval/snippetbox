package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// home is a handler function which writes a byte slice as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Check for an exact URL path match and if no match; send a 404 response to the client.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// showSnippet displays specific snippets.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and convert to an integer; otherwise respond w/ 404.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Interpolate the id value with the response and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with id: %d", id)
}

// createSnippet displays a form to submit new snippets.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Check whether the request is using POST or not. If not, send a 405 status code and a response body message.
	if r.Method != "POST" {
		// Add 'Allow: POST' to response header map.
		w.Header().Set("Allow", "POST")
		// Use the http.Error() function that uses w.WriteHeader() and w.Write() under the hood to send a string as the response body.
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Initialize a new servemux, then register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Start a new web server and pass the TCP network address to listen on and the servemux just created.
	// If http.ListenAndServe() returns an error, an error will be thrown.
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
