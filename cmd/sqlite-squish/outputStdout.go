package main

import (
	"fmt"

	"github.com/griffithsh/sqlite-squish/database"
)

func outputStdout(d *database.Database) error {
	fmt.Println(d.String())
	return nil
}
