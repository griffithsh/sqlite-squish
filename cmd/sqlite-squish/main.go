package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/griffithsh/sql-squish/database"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var concatted string
	for scanner.Scan() {
		concatted = concatted + scanner.Text()
	}
	d, err := database.FromString(concatted)
	if err != nil {
		log.Fatal(err)
	}
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
}
