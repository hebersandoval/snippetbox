package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// home is a handler function which writes a byte slice as the response body.
// Now is a method against *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check for an exact URL path match and if no match; send a 404 response to the client.
	if r.URL.Path != "/" {
		app.notFound(w) // Use the notFound() helper
		return
	}
	// Initialize a slice containing the paths to files. Note: home.page.tmpl file must be *first* in the slice.
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	// Read the template files into a template set. If there's an error, log the details.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Method against application can access its fields.
		app.serverError(w, err) // Use the serverError() helper.
		return
	}
	// Write the template's content as the response body on the template set and send any dynamic data.
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err) // Use the serverError() helper.
	}
}

// showSnippet displays specific snippets.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and convert to an integer; otherwise respond w/ 404.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Use the notFound() helper.
		return
	}
	// Interpolate the id value with the response and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with id: %d", id)
}

// createSnippet displays a form to submit new snippets.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Check whether the request is using POST or not. If not, send a 405 status code and a response body message.
	if r.Method != "POST" {
		// Add 'Allow: POST' to response header map.
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
