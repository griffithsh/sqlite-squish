package statement

import (
	"fmt"
	"regexp"
	"strings"
)

// Statement captures a single sql statement, ie INSERT INTO (...) VALUES (...);
type Statement struct {
	SQL  string
	Verb Verb // A statement type; INSERT, CREATE, etc
}

// Verb describes a type of sql statement, INSERT, CREATE, etc
type Verb int

const (
	// Create represents a Create Table statement
	Create Verb = iota
	// Insert represents an Insert statement
	Insert
)

func (v Verb) String() string {
	switch v {
	case Create:
		return "Create"
	case Insert:
		return "Insert"
	default:
		return ""
	}
}

// Table returns the name of the table that this statement operates upon.
func (s *Statement) Table() string {
	var reg *regexp.Regexp
	switch s.Verb {
	case Create:
		reg = regexp.MustCompile(`(?i)^CREATE TABLE \[?(?P<X>[A-Za-z\-_0-9]+)\]?`)
	case Insert:
		reg = regexp.MustCompile(`(?i)^INSERT INTO "?(?P<X>[A-Za-z\-_0-9]+)"?`)
	default:
		return ""
	}
	matches := reg.FindStringSubmatch(s.SQL)
	if len(matches) == 0 {
		return ""
	}
	// The first match should be *everything* from the regex, and the second
	// match should be only the capture group.
	return matches[1]
}

// Dependencies determines the foreign key dependencies of the table if the
// statement is a Create Statement
func (s *Statement) Dependencies() []string {
	if s.Verb != Create {
		return []string{}
	}

	reg := regexp.MustCompile(`(?i)REFERENCES \[?"?(?P<X>[A-Z\-_0-9]+)`)
	matches := reg.FindAllStringSubmatch(s.SQL, -1)
	var result []string
	for _, match := range matches {
		if len(match) == 0 {
			continue
		}
		if s.Table() != match[1] {
			result = append(result, match[1])
		}
	}

	return result
}

// Statements type is needed to satisfy the sort interface
type Statements []Statement

func (slice Statements) Len() int {
	return len(slice)
}

func (slice Statements) Less(i, j int) bool {
	return slice[i].Verb < slice[j].Verb
}

func (slice Statements) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// FromString constructs a Statement from a string
func FromString(s string) (*Statement, error) {
	s = strings.Trim(s, "\r\n\t ")
	if isCreate, err := regexp.MatchString("(?i)^CREATE TABLE ", s); isCreate && err == nil {
		return &Statement{SQL: s, Verb: Create}, nil
	}
	if isInsert, err := regexp.MatchString("(?i)^INSERT INTO ", s); isInsert && err == nil {
		return &Statement{SQL: s, Verb: Insert}, nil
	}
	return nil, fmt.Errorf("A Statement could not be constructed from \"%s\"", s)
}

func (s *Statement) String() string {
	return s.SQL
}
