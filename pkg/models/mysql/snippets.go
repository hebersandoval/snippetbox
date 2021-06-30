package mysql

import (
	"database/sql"
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
	return nil, nil
}

// Latest return 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}