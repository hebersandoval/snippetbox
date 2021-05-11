package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define a new command-line flag, a default value and description. The value will be stored at runtime.
	addr := flag.String("addr", ":8080", "HTTP network address.")

	// Parse the command-line and read in the flag value and assign it to the "addr" variable. Should be called before using "addr".
	flag.Parse()

	// Create a logger for writing information and error messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new servemux, then register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create a file server which serves files out of the "./ui/static" directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register the file server as the handler for all URL paths that start with "/static/" and strip the prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Initialize a new http.Server struct. Now the server can use the custom errorLog in the event of any problems.
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	// Start a new web server and pass the TCP network address to listen on and the servemux just created.
	// If http.ListenAndServe() returns an error, an error will be thrown.
	infoLog.Printf("Starting server on %s", *addr) // Deference the pointer returned from flag.String()
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}


// Note: To run server from Windows OS -> go run ./cmd/web/.