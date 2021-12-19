package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError helper writes an error message and stack trace to the errorLog then sends a
// generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace) // Report one step back in the stack trace.
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError helper sends a specific status code and corresponding description to the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	// Use the http.Error() function that uses w.WriteHeader() and w.Write() under the hood to send a string as the response body.
	http.Error(w, http.StatusText(status), status)
}

// notFound helper is a convenient wrapper around clientError which sends a "404 Not Found" response to the use.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	// Retrieve the appropiate template set from the cache based on the page name (like 'home.page.tmpl').
	// If no entry exists in the cache with the provided name, call the serverError helper method.
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	// Execute the template set, passing in any dynamic data.
	err := ts.Execute(w, td)
	if err != nil {
		app.serverError(w, err)
	}
}
