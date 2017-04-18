package main

import (
	"io/ioutil"

	"fmt"

	"os"

	"github.com/griffithsh/sqlite-squish/database"
)

func outputText(d *database.Database, dir string) error {
	// If dir does not exist, create it.
	info, err := os.Lstat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}
			info, err = os.Lstat(dir)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Else if dir is something that isnt a directory, error out.
	if !info.IsDir() {
		return fmt.Errorf("cannot output to non-directory %s", dir)
	}

	tableNames, err := d.SortedTables()
	if err != nil {
		return err
	}
	for _, name := range tableNames {
		file := fmt.Sprintf("%s\n", d.Tables[name].String())
		filepath := fmt.Sprintf("%s/%s.sql", info.Name(), name)
		err = ioutil.WriteFile(filepath, []byte(file), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
