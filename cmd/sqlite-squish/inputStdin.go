package main

import (
	"bufio"
	"os"
)

func inputStdin() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	var concatted string
	for scanner.Scan() {
		concatted = concatted + scanner.Text()
	}
	return concatted, nil
}
