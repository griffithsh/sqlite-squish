package main

import (
	"fmt"

	"github.com/griffithsh/sql-squish/database"
)

func outputStdout(d database.Database) error {
	for _, t := range d.Tables {
		fmt.Println(t.String())
	}
	// files, err := db.AsSQL()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for file := range files {
	// 	fmt.Print(file)
	// }
	return nil
}
