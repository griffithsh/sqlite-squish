package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/griffithsh/sqlite-squish/database/from"
)

func TestDirToSqliteToText(t *testing.T) {
	err := os.Mkdir("./testdata/output", 0755)
	if err != nil {
		t.Fatalf("make test output dir: %s", err)
	}
	defer os.RemoveAll("./testdata/output")

	// Start by interpreting the testdata.
	d, err := from.Directory("./testdata")
	if err != nil {
		t.Fatalf("input testdata: %s", err)
	}

	// Write the interpretation to a tmp sqlite file.
	err = d.File("./testdata/output/tmp.sqlite")
	if err != nil {
		t.Fatalf("output test db: %s", err)
	}

	// Interpret the tmp sqlite database we just wrote.
	d, err = from.SQLite("./testdata/output/tmp.sqlite")
	if err != nil {
		t.Fatalf("input temp db: %s", err)
	}

	// Write the freshly interpreted database to the testdata output directory.
	err = d.Directory("./testdata/output")
	if err != nil {
		t.Fatalf("output test data: %s", err)
	}

	// Compare the original testdata to the output testdata.
	files, err := ioutil.ReadDir("./testdata")
	for _, file := range files {
		name := file.Name()
		if file.IsDir() {
			continue
		}
		// FIXME(griffithsh): This is unpalatable to users without `diff` on their PATH
		cmd := exec.Command("diff", "-u", "./testdata/"+name, "./testdata/output/"+name)

		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out

		if err := cmd.Run(); err != nil {
			t.Errorf("%s\n%s", err, out.String())
		}
	}
}
