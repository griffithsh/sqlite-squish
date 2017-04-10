package database

import (
	"fmt"
	"strings"

	"log"

	"github.com/griffithsh/sql-squish/statement"
	"github.com/griffithsh/sql-squish/table"
)

// Database is a collection of the tables in a database
type Database struct {
	Tables map[string]*table.Table
}

func (d Database) String() string {
	// sort tables by their dependencies
	return "todo - the string representation of this database"
}

// AsSQL returns a collection of strings each representing the file contents of
// a plaintext sql representation of a Table in this Database.
func (d Database) AsSQL() ([]string, error) {
	// TODO sort tables by their dependency graphs
	// return a string representation of each sorted Table
	return []string{}, nil
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
		stmt, err := statement.FromString(fmt.Sprintf("%s;", str))
		if err != nil {
			fmt.Println(err)
			continue
		}
		t, ok := output.Tables[stmt.Table()]
		if !ok {
			output.Tables[stmt.Table()] = &table.Table{Name: stmt.Table()}
			t = output.Tables[stmt.Table()]
			//fmt.Printf("output now has %d tables\n", len(output.Tables))
		}
		err = t.Add(stmt)
		//fmt.Printf("table %s now has %d statements\n", t.Name, len(t.Statements))
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
