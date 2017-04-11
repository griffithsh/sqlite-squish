package main

import (
	"fmt"

	"github.com/griffithsh/sql-squish/database"
)

func outputStdout(d database.Database) error {
	files, err := d.SortedTables()
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Print(file)
	}
	return nil
}
