package from

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/griffithsh/sqlite-squish/database"
)

// Directory composes a Database from a directory containing .sql files.
func Directory(dir string) (*database.Database, error) {
	info, err := os.Lstat(dir)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("Cannot read from non-directory %s", dir)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var output string
	for _, file := range files {
		// if file ends in ".sql" ...
		if isSQLFile(file.Name()) {
			contents, err := ioutil.ReadFile(dir + "/" + file.Name())
			if err != nil {
				return nil, err
			}
			output = output + string(contents)
		}
	}
	return database.FromString(output)
}

func isSQLFile(fileName string) bool {
	return strings.ToUpper(filepath.Ext(fileName)) == ".SQL"
}
