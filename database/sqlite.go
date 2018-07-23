package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// File outputs the Database as a binary sqlite file.
func (d *Database) File(file string) error {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return fmt.Errorf("create db file: %s", err)
	}
	for _, s := range strings.Split(d.String(), "\n") {
		_, err := db.Exec(s)
		if err != nil {
			return fmt.Errorf("execute sql %s: %s", s, err)
		}
	}
	return nil
}
