package table

import (
	"errors"
	"sort"

	"github.com/griffithsh/sql-squish/statement"
)

// Table is a collection of Statements, and a representation of a sqlite
// database table. It should have a way to input from and output to both a
// plaintext "Decomposition" of the table, and a binary sqlite database file.
type Table struct {
	Statements statement.Statements
	Name       string

	// If a table has foreign keys, then it has Dependencies on those other tables.
	Dependencies []string
}

// Add a statement to the Table
func (t *Table) Add(s statement.Statement) error {
	// It's not coherent to have multiple CREATE statements for a single table.
	if len(t.Statements) > 0 && s.Verb == statement.Create && t.Statements[0].Verb == statement.Create {
		return errors.New("Cannot add multiple Create Statements to Table")
	}

	// If we are adding a new create statement to this Table, then we need to
	// determine and store the Table's dependencies.
	if s.Verb == statement.Create {
		t.Dependencies = s.Dependencies()
	}

	// After adding a new Statement, it's imperitive to keep the Statements sorted
	t.Statements = append(t.Statements, s)
	sort.Sort(t.Statements)

	return nil
}
