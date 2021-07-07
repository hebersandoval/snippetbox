package mysql

import (
	"database/sql"
	"errors"
	"github.com/hebersandoval/snippetbox/pkg/models"
)

// SnippetModel type which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert inserts new snippet into db.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
		VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	// Use the Exec() method on the embedded connection pool to execute the statement.
	// This method returns a sql.Result object.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	// Use the LastInsert() method on the result object to get the ID of our newly inserted record in the snippet table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	// The ID returned has the type int64, so we convert it to an int type before returning.
	return int(id), nil
}

// Get return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error){
	stmt := `SELECT id, title, content, created, expires FROM snippets
		WHERE expires > UTC_TIMESTAMP() AND id = ?`
	// Passing in the untrusted id variable as the value fo the placeholder parameter.
	// This returns a pointer to a sql.Row object which holds the result from the db.
	row := m.DB.QueryRow(stmt, id)
	// Initialize a pointer to a new zeroed Snippet struct.
	s := &models.Snippet{}
	// Use row.Scan() to copy the values from each field in sql.Row to the corresponding field in the Snippet struct.
	// Arguments to row.Scan are "pointers" to the place you want to copy the data into; numbers of args must be
	// exactly the same the number of columns returned by the statement.
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// If the query returns no rows, then row.Scan() will return a sql.ErrNoRows error.
		// Use errors.Is() function to check for that error specifically, and return our own models.ErrNoRecord error instead.
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	// If everything went OK then return the Snippet object.
	return s, nil
}

// Latest return 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
		WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	// Use the Query() method on the connection pool to execute our SQL statement, which returns a sql.Rows resultset
	// containing the result of our query.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// defer should be after checking for an error from Query() method, otherwise, if error is returned, you'll
	// get a panic due to trying to close a nil resultset.
	defer rows.Close()
	// Initialize an empty slice to hold the models.Snippets objects.
	snippets := []*models.Snippet{}
	// Use rows.Next() to iterate through the rows in the resultset. After iteration is completed, then the resultset
	// automatically closes itself and frees-up the underlying db connection.
	for rows.Next() {
		// Create a pointer to a new zeroed Snippet struct.
		s := &models.Snippet{}
		// Use rows.Scan() to copy the values from each field in the row to the new Snippet object that was created.
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		snippets = append(snippets, s)
	}
	// When the rows.Next() loop has finished, call rows.Err() to retrieve any error that was encountered during iteration.
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// If all went well, then return the Snippets slice.
	return snippets, nil
}