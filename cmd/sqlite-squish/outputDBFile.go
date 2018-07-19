package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/griffithsh/sqlite-squish/database"
	_ "github.com/mattn/go-sqlite3"
)

func outputDBFile(d *database.Database, file string) error {
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
