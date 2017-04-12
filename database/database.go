package database

import (
	"fmt"
	"log"
	"strings"

	"github.com/griffithsh/sqlite-squish/statement"
	"github.com/griffithsh/sqlite-squish/table"
)

// Database is a collection of the tables in a database
type Database struct {
	Tables map[string]*table.Table
}

func (d Database) String() string {
	tables, err := d.SortedTables()
	if err != nil {
		return err.Error()
	}
	return strings.Join(tables, "\n")
}

// SortedTables returns a collection of strings, each representing the file
// contents of a plaintext sql representation of a Table in this Database.
// The tables are sorted such that tables that are referenced by other tables
// with foreign keys are output before their dependants.
func (d Database) SortedTables() ([]string, error) {
	// TODO - dependency sorting
	var tables []string
	for _, t := range d.Tables {
		tables = append(tables, t.String())
	}
	return tables, nil
}

// FromString creates a logical representation of a sqlite database from a
// series of sql statements in a string.
func FromString(input string) (Database, error) {
	// NB, expecting input from a command like this:
	// (sqlite3 db.sqlite .schema; sqlite3 db.sqlite .dump | grep '^INSERT ') | ...
	// or even
	// sqlite3 db.sqlite .dump | ...

	strs := splitStatements(input)

	var output Database
	output.Tables = map[string]*table.Table{}

	for _, str := range strs {
		stmt, err := statement.FromString(fmt.Sprintf("%s", str))
		if err != nil {
			log.Println(err)
			continue
		}
		t, ok := output.Tables[stmt.Table()]
		if !ok {
			output.Tables[stmt.Table()] = &table.Table{Name: stmt.Table()}
			t = output.Tables[stmt.Table()]
		}
		err = t.Add(stmt)
		if err != nil {
			log.Print(err)
		}
	}

	return output, nil
}

func splitStatements(input string) []string {
	stmts := strings.Split(input, ";")

	// If there is a VARCHAR with a semicolon in it, the naive Split will cut a
	// statement in half. We need to correct any such bad splits if they occur.
	var corrected []string

	for _, stmt := range stmts {
		// Given a collection of statements all ending in a semicolon, then
		// when Split(), an empty string should be the last element...
		if stmt == "" {
			continue
		}

		// We need to re-add this semicolon on to the end of each split to
		// return the statements to their original forms, as it has been
		// removed by strings.Split().
		stmt = stmt + ";"

		// If the last element in corrected has an uneven number of
		// apostrophes, append this stmt to the last string in corrected as a
		// semicolon inside a string has been found, and the string is now
		// missing it's closing apostrophe.
		size := len(corrected)
		if size > 0 && strings.Count(corrected[size-1], "'")%2 == 1 {
			corrected[size-1] = corrected[size-1] + stmt
		} else {
			// Otherwise append a new string to corrected.
			corrected = append(corrected, stmt)
		}
	}
	return corrected
}
