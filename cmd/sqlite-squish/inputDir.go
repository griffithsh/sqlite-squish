package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func inputDir(dir string) (string, error) {
	info, err := os.Lstat(dir)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", fmt.Errorf("Cannot read from non-directory %s", dir)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}
	var output string
	for _, file := range files {
		// if file ends in ".sql" ...
		if isSQLFile(file.Name()) {
			contents, err := ioutil.ReadFile(dir + "/" + file.Name())
			if err != nil {
				return "", err
			}
			output = output + string(contents)
		}
	}
	return output, nil
}

func isSQLFile(fileName string) bool {
	return strings.ToUpper(filepath.Ext(fileName)) == ".SQL"
}
