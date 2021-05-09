package main

import (
	"fmt"
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


// Note: To run server from Windows OS -> go run ./cmd/web/.