package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Hold application-wide dependencies for web app.
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	// Define a new command-line flag, a default value and description. The value will be stored at runtime.
	addr := flag.String("addr", ":8080", "HTTP network address.")

	// Parse the command-line and read in the flag value and assign it to the "addr" variable. Should be called before using "addr".
	flag.Parse()

	// Create a logger for writing information and error messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	// Initialize a new http.Server struct. Now the server can use the custom errorLog in the event of any problems.
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(), // Call the new app.routes() method.
	}

	// Start a new web server and pass the TCP network address to listen on and the servemux just created.
	// If http.ListenAndServe() returns an error, an error will be thrown.
	infoLog.Printf("Starting server on %s", *addr) // Deference the pointer returned from flag.String()
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}


// Note: To run server from Windows OS -> go run ./cmd/web/.