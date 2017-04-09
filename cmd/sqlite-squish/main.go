package main

import (
	"log"

	"github.com/griffithsh/sql-squish/database"
)

func main() {
	concatted, err := inputStdin()

	if err != nil {
		log.Fatal(err)
	}
	d, err := database.FromString(concatted)
	if err != nil {
		log.Fatal(err)
	}

	outputStdout(d)
}
