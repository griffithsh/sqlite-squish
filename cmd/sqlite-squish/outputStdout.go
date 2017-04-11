package main

import (
	"fmt"

	"github.com/griffithsh/sql-squish/database"
)

func outputStdout(d database.Database) error {
	fmt.Println(d.String())
	return nil
}
