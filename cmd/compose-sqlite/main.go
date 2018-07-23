/*
program compose-sqlite takes a directory of .sql text files and outputs a
binary sqlite database file.

Example:

	compose-sqlite data/

This will compose data.sqlite3 from the *.sql files in the directory anmed
data. A .sqlite3 file extension is added to the composed database.

If there is a .sql file extension on the directory, then it is ignored. When
there is a data.sql directory present, then the below will also compose it into data.sqlite3.

	compose.sqlite data
*/
package main

import (
	"os"
	"path"
	"strings"

	"github.com/griffithsh/sqlite-squish/database/from"
)

func main() {
	dir := os.Args[1]
	d, err := from.Directory(dir)
	if err != nil {
		d, err = from.Directory(dir + ".sql")
		if err != nil {
			panic(err)
		}
	}
	name := strings.TrimSuffix(path.Base(dir), ".sql") + ".sqlite3"
	os.Remove(name)
	if err = d.File(name); err != nil {
		panic(err)
	}
}
