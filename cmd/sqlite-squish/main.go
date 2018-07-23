package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"io/ioutil"

	"github.com/griffithsh/sqlite-squish/database"
	"github.com/griffithsh/sqlite-squish/database/from"
)

var (
	inFlag      = flag.String("in", "", "Specifies a directory of .sql files, or a sqlite database file to use as input")
	outDBFlag   = flag.String("out-db", "", "The input will be written into a sqlite database file with this filename")
	outDirFlag  = flag.String("out-dir", "", "The input will be written into this directory as a a collection of .sql files, one file per table, where the filename is constructed from the table name")
	verboseFlag = flag.Bool("v", false, "Verbose output")
)

func main() {
	flag.Parse()

	if !*verboseFlag {
		log.SetOutput(ioutil.Discard)
	}

	var d *database.Database
	var err error

	// If there is nothing provided for the -in flag, read from Stdin.
	if *inFlag == "" {
		d, err = from.Reader(os.Stdin)
		exitOnError(err)
	} else {
		info, err := os.Stat(*inFlag)
		exitOnError(err)
		if info.IsDir() {
			d, err = from.Directory(*inFlag)
			exitOnError(err)
		} else {
			d, err = from.SQLite(*inFlag)
			exitOnError(err)
		}
	}

	// This is the point at which the input, wherever it came from, is
	// composed into the intermediary format, ready for output.
	// d, err := database.FromString(input)
	// exitOnError(err)

	// The flow should be that if the user specifies a text output, or a db
	// output, then write those output(s). If they specify neither, then just
	// fallback to Stdout.
	if *outDirFlag == "" && *outDBFlag == "" {
		fmt.Println(d.String())
		return
	}
	if *outDirFlag != "" {
		err = d.Directory(*outDirFlag)
		exitOnError(err)
	}
	if *outDBFlag != "" {
		err = d.File(*outDBFlag)
		exitOnError(err)
	}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
