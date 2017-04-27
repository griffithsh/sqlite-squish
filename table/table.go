package table

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/griffithsh/sqlite-squish/statement"
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
func (t *Table) Add(s *statement.Statement) error {
	// It's not coherent to have multiple CREATE statements for a single table.
	if len(t.Statements) > 0 && s.Verb == statement.Create && t.Statements[0].Verb == statement.Create {
		return fmt.Errorf("Cannot add multiple Create Statements to Table %s", t.Name)
	}

	// It's a mistake to add a statement for another table
	if t.Name != s.Table() {
		return fmt.Errorf("Cannot add Statement for %s to table %s", s.Table(), t.Name)
	}

	// If we are adding a new create statement to this Table, then we need to
	// determine and store the Table's dependencies.
	if s.Verb == statement.Create {
		t.Dependencies = s.Dependencies()
	}

	// After adding a new Statement, it's imperative to keep the Statements sorted.
	t.Statements = append(t.Statements, *s)
	sort.Sort(t.Statements)

	return nil
}

type keyedStatement struct {
	statement.Statement
	key string
}
type keyedStatements []keyedStatement

func (slice keyedStatements) Len() int {
	return len(slice)
}

func (slice keyedStatements) Less(i, j int) bool {
	if slice[i].Verb == slice[j].Verb {
		// try integer cast ...
		iKey, iErr := strconv.ParseInt(slice[i].key, 10, 64)
		jKey, jErr := strconv.ParseInt(slice[j].key, 10, 64)
		if iErr != nil || jErr != nil {
			return slice[i].key < slice[j].key
		}
		return iKey < jKey
	}
	return slice[i].Verb < slice[j].Verb
}

func (slice keyedStatements) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (t *Table) String() string {
	if len(t.Statements) == 0 {
		return ""
	}

	// Result now contains the CREATE TABLE statement.
	result := t.Statements[0].String()

	// The assumption here is that the primary key of the table is the first
	// column, so if we pull it out of the SQL, we can order by that value...
	var inserts keyedStatements
	reg := regexp.MustCompile(`(?i)\('?(?P<X>[A-Z\-_0-9]+)'?[\,\)]`)
	for _, s := range t.Statements[1:] {
		matches := reg.FindStringSubmatch(s.SQL)
		var key string
		if len(matches) > 1 {
			key = matches[1]
		}
		inserts = append(inserts, keyedStatement{s, key})
	}
	sort.Sort(inserts)

	for _, s := range inserts {
		result = fmt.Sprintf("%s\n%s", result, s.String())
	}

	// We want to blank out the sqlite_sequence table because it accumulates
	// rows if you don't.
	if t.Name == "sqlite_sequence" {
		result = fmt.Sprintf("DELETE FROM sqlite_sequence;\n%s", result)
	}
	return result
}
