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
