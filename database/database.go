package database

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"sort"

	"github.com/griffithsh/sqlite-squish/statement"
	"github.com/griffithsh/sqlite-squish/table"
)

// Database is a collection of the tables in a database
type Database struct {
	Tables map[string]*table.Table
}

func (d Database) String() string {
	tableNames, err := d.SortedTables()
	if err != nil {
		return err.Error()
	}
	var tables string
	tables = "BEGIN TRANSACTION;\n"
	for _, name := range tableNames {
		tables = fmt.Sprintf("%s%s\n", tables, d.Tables[name].String())
	}
	tables = fmt.Sprintf("%s%s\n", tables, "COMMIT;")
	return tables
}

// SortedTables returns the table names of the database, sorted alphabetically
// first, but also sorting tables with references to other tables after those
// tables they refer to.
func (d Database) SortedTables() ([]string, error) {
	var sorted []table.Table
	var visited []table.Table
	var err error

	// This iteration through alphabetically sorted keys approach is so that
	// the output is deterministic. Tables should be alphabetically sorted when
	// there are multiple valid sorts with regards to dependencies.
	var keys []string
	for _, t := range d.Tables {
		keys = append(keys, t.Name)
	}
	sort.Strings(keys)

	for _, key := range keys {
		t := d.Tables[key]
		sorted, visited, err = visit(t, visited, sorted, d.dependencies)
		if err != nil {
			return nil, err
		}
	}

	// Convert sorted tables to their table name as a string.
	var names []string
	for _, n := range sorted {
		names = append(names, n.Name)
	}
	return names, nil
}

func (d Database) dependencies(t *table.Table) []*table.Table {
	var found []*table.Table

	// Sort dependencies so that output of SortedTables is deterministic.
	deps := t.Dependencies
	sort.Strings(deps)

	for _, sdep := range deps {
		if d.Tables[sdep] == nil {
			continue
		}
		found = append(found, d.Tables[sdep])
	}

	return found
}

func visit(t *table.Table, visited []table.Table, sorted []table.Table, dependencies func(*table.Table) []*table.Table) ([]table.Table, []table.Table, error) {
	if !contains(visited, t) {
		visited = append(visited, *t)
		for _, dep := range dependencies(t) {
			var err error
			sorted, visited, err = visit(dep, visited, sorted, dependencies)
			if err != nil {
				return nil, nil, err
			}
		}
		sorted = append(sorted, *t)
	} else if !contains(sorted, t) {
		return nil, nil, errors.New("Cyclic dependency error")
	}
	return sorted, visited, nil
}

func contains(s []table.Table, find *table.Table) bool {
	for _, t := range s {
		if find == nil {
			fmt.Printf("panic on nil ptr deref: %s\n", t.Name)
		}
		if t.Name == find.Name {
			return true
		}
	}
	return false
}

// FromString creates a logical representation of a sqlite database from a
// series of sql statements in a string.
func FromString(input string) (*Database, error) {
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

	return &output, nil
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
