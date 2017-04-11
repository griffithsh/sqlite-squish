package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"io/ioutil"

	"github.com/griffithsh/sql-squish/database"
)

var (
	inFlag      = flag.String("in", "", "Specifies a directory of .sql files, or a sqlite database file to use as input")
	outDBFlag   = flag.String("out-db", "", "The input will be written into a sqlite database file with this filename")
	outTextFlag = flag.String("out-text", "", "The input will be written into this directory as a a collection of .sql files, one file per table, where the filename is constructed from the table name")
	verboseFlag = flag.Bool("v", false, "Verbose output")
)

func main() {
	flag.Parse()

	if !*verboseFlag {
		log.SetOutput(ioutil.Discard)
	}

	var input string
	var err error

	// if there is nothing provided for the -in flag, read from Stdin
	if *inFlag == "" {
		input, err = inputStdin()
		exitOnError(err)
	} else {
		info, err := os.Stat(*inFlag)
		exitOnError(err)
		if info.IsDir() {
			input, err = inputDir(*inFlag)
			exitOnError(err)
		} else {
			input, err = inputFile(*inFlag)
			exitOnError(err)
		}
	}

	// This is the point at which the input, wherever it came from, is
	// composed into the intermediary format, ready for output.
	d, err := database.FromString(input)
	exitOnError(err)

	// The flow should be that if the user specifies a text output, or a db
	// output, then write those output(s). If they specify neither, then just
	// falback to stdout.
	if *outTextFlag == "" && *outDBFlag == "" {
		err = outputStdout(d)
		exitOnError(err)
		return
	}
	if *outTextFlag != "" {
		err = outputText(*outTextFlag)
		exitOnError(err)
	}
	if *outDBFlag != "" {
		err = outputDB(*outDBFlag)
		exitOnError(err)
	}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
