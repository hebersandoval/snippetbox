package main

import (
	"database/sql"
	"flag"
	"github.com/hebersandoval/snippetbox/pkg/models/mysql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Hold application-wide dependencies for web app.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {
	// Define a new command-line flag, a default value and description. The value will be stored at runtime.
	addr := flag.String("addr", ":8080", "HTTP network address.")

	// Define a new command-line flag for the MySQL DSN string.
	dsn := flag.String("dsn", "web:secret@/snippetbox?parseTime=true", "MySQL data source name")

	// Parse the command-line and read in the flag value and assign it to the "addr" and "dsn" variables.
	// Should be called before using "addr" and "dsn".
	flag.Parse()

	// Create a logger for writing information and error messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Creating a connection pool into separate openDB() below, passing the DSN from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// defer call to db.Close() so that the connection pool is closed before the main() exits.
	defer db.Close()

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	// Initialize a new http.Server struct. Now the server can use the custom errorLog in the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // Call the new app.routes() method.
	}

	// Start a new web server and pass the TCP network address to listen on and the servemux just created.
	// If http.ListenAndServe() returns an error, an error will be thrown.
	infoLog.Printf("Starting server on %s", *addr) // Deference the pointer returned from flag.String()
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// openDB() wraps sql.Open() and returns a sql.DB connection pool for a given a DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Note: To run server from Windows OS -> go run ./cmd/web/.
