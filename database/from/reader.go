package from

import (
	"bufio"
	"io"

	"github.com/griffithsh/sqlite-squish/database"
)

// Reader composes a Database from an io.Reader, such as os.Stdin
func Reader(r io.Reader) (*database.Database, error) {
	scanner := bufio.NewScanner(r)

	var concatted string
	for scanner.Scan() {
		concatted = concatted + scanner.Text()
	}
	return database.FromString(concatted)
}
