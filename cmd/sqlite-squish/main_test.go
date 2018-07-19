package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/griffithsh/sqlite-squish/database"
)

func TestDirToSqliteToText(t *testing.T) {
	err := os.Mkdir("./testdata/output", 0755)
	if err != nil {
		t.Fatalf("make test output dir: %s", err)
	}
	defer os.RemoveAll("./testdata/output")

	in, err := inputDir("./testdata")
	if err != nil {
		t.Fatalf("input testdata: %s", err)
	}

	d, err := database.FromString(in)
	if err != nil {
		t.Fatalf("interpret testdata: %s", err)
	}

	err = outputDBFile(&d, "./testdata/output/tmp.sqlite")
	if err != nil {
		t.Fatalf("output test db: %s", err)
	}
	in, err = inputDBFile("./testdata/output/tmp.sqlite")
	if err != nil {
		t.Fatalf("input temp db: %s", err)
	}

	d, err = database.FromString(in)
	if err != nil {
		t.Fatalf("interpret temp db: %s", err)
	}

	err = outputText(&d, "./testdata/output")
	if err != nil {
		t.Fatalf("output test data: %s", err)
	}

	// diff
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
