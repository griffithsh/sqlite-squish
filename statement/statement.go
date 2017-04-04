package statement

// Verb describes a type of sql statement, INSERT, CREATE, etc
type Verb int

const (
	// Insert represents an Insert statement
	Insert Verb = iota
	// Create represents a Create Table statement
	Create
)

// Statement captures a single sql statement, ie INSERT INTO (...) VALUES (...);
type Statement struct {
	SQL  string
	Verb Verb // A statement type; INSERT, CREATE, etc
}
