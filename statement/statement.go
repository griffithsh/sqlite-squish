package statement

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

// Dependencies determines the foreign key dependencies of the table if the
// statement is a Create Statement
func (s *Statement) Dependencies() []string {
	if s.Verb != Create {
		return []string{}
	}
	// FIXME I feel like it's a bit premature to work on the dependency
	// resolver, so for now, we'll return some junk dependencies...
	return []string{"A", "B", "C"}
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
