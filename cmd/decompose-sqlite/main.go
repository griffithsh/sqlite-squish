/*
program decompose-sqlite takes a binary sqlite database file and decomposes
it into a set of text *.sql files

Example:

	decompose-sqlite database.sqlite

creates a file for every table in database.sqlite under database.sql/ that
contains a CREATE TABLE statement and zero or more INSERT INTO statements.

If a file extension is not provided, then .sqlite and .sqlite3 are also
tried. The below examples are equivalent when a database called data.sqlite
is present.

	decompose-sqlite data
*/
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/griffithsh/sqlite-squish/database/from"
)

func good(file string) bool {
	fi, err := os.Stat(file)
	return err == nil && fi != nil && !fi.IsDir()
}

func inferName(file string) (string, error) {
	try := file
	if good(try) {
		return try, nil
	}

	try = file + ".sqlite"
	if good(try) {
		return try, nil
	}

	try = file + ".sqlite3"
	if good(try) {
		return try, nil
	}

	try = strings.TrimSuffix(file, ".sql") + ".sqlite"
	if good(try) {
		return try, nil
	}

	try = strings.TrimSuffix(file, ".sql") + ".sqlite3"
	if good(try) {
		return try, nil
	}

	return "", fmt.Errorf("could not infer database name from %s", file)
}

func inferOut(name string) string {
	name = strings.TrimSuffix(name, ".sqlite3")
	name = strings.TrimSuffix(name, ".sqlite")
	name = name + ".sql"
	return name
}

func main() {
	input := os.Args[1]

	name, err := inferName(input)
	if err != nil {
		panic(err)
	}

	db, err := from.SQLite(name)
	if err != nil {
		panic(err)
	}

	output := inferOut(name)
	err = os.RemoveAll(output)
	if err != nil {
		panic(err)
	}
	err = db.Directory(output)
	if err != nil {
		panic(err)
	}
}
