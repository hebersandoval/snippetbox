package main

import (
	"errors"
	"fmt"
	"github.com/hebersandoval/snippetbox/pkg/models"
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
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Use the render helper.
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: snippets,
	})

	/*
		// Create an instance of a templateData struct holding the slice of snippets.
		data := &templateData{Snippets: snippets}

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
		// Pass in the templateData struct when executing the template.
		err = ts.Execute(w, data)
		if err != nil {
			app.serverError(w, err) // Use the serverError() helper.
		}
	*/
}

// showSnippet displays specific snippets.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and convert to an integer; otherwise respond w/ 404.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Use the notFound() helper.
		return
	}
	// Use the SnippetModel object's Get Method to retrieve the data for a specific record based on its ID.
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Use the render helper.
	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: snippet,
	})

	/*
		// Create an instance of a templateData struct holding the snippet data.
		data := &templateData{Snippet: snippet}
		// Initialize a slice containing the paths to the show.page.tmpl file, plus the base layout and footer partial.
		files := []string{
			"./ui/html/show.page.tmpl",
			"./ui/html/base.layout.tmpl",
			"./ui/html/footer.partial.tmpl",
		}
		// Parse the template files...
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.serverError(w, err)
			return
		}
		// ...and then execute them; pass the templateData struct, which contains a model.Snippet field.
		err = ts.Execute(w, data)
		if err != nil {
			app.serverError(w, err)
		}
		// Interpolate the id value with the response and write it to the http.ResponseWriter.
		//fmt.Fprintf(w, "Display a specific snippet with id: %v", s)
	*/
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
	// Dummy data
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := "7"
	// Pass the data to the SnippetModel.Insert() method receiving the ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
