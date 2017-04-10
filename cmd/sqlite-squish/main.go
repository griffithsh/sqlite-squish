package main

import (
	"flag"
	"log"
	"os"

	"github.com/griffithsh/sql-squish/database"
)

var (
	inFlag      = flag.String("in", "", "Specifies a directory of .sql files, or a sqlite database file to use as input")
	outDBFlag   = flag.String("out-db", "", "The input will be written into a sqlite database file with this filename")
	outTextFlag = flag.String("out-text", "", "The input will be written into this directory as a a collection of .sql files, one file per table, where the filename is constructed from the table name")
)

func main() {
	flag.Parse()

	var input string
	var err error

	// if there is nothing provided for the -in flag, read from Stdin
	if *inFlag == "" {
		input, err = inputStdin()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		info, err := os.Stat(*inFlag)
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() {
			input, err = inputDir(*inFlag)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			input, err = inputFile(*inFlag)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// This is the point at which the input, wherever it came from, is
	// composed into the intermediary format, ready for output.
	d, err := database.FromString(input)
	if err != nil {
		log.Fatal(err)
	}

	// The flow should be that if the user specifies a text output, or a db
	// output, then write those output(s). If they specify neither, then just
	// falback to stdout.
	if *outTextFlag == "" && *outDBFlag == "" {
		err = outputStdout(d)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if *outTextFlag != "" {
		err = outputText(*outTextFlag)
		if err != nil {
			log.Fatal(err)
		}
	}
	if *outDBFlag != "" {
		err = outputDB(*outDBFlag)
		if err != nil {
			log.Fatal(err)
		}
	}
}
