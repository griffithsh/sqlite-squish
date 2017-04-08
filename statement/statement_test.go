package statement

import (
	"sort"
	"testing"
)

func TestSortCreateBeforeInsert(t *testing.T) {
	s := Statements{
		{Verb: Insert},
		{Verb: Create},
	}

	sort.Sort(s)

	if s[0].Verb != Create {
		t.Errorf("first element should have been %s but was %s", Create, s[0].Verb)
	}
	if s[1].Verb != Insert {
		t.Errorf("second element should have been %s but was %s", Insert, s[0].Verb)
	}
}

func TestStatementFromString(t *testing.T) {
	tests := []struct {
		input       string
		expected    *Statement
		expectError bool
	}{
		{input: "Create Table [Bananas] ([Id] INTEGER);", expected: &Statement{Verb: Create, SQL: "Create Table [Bananas] ([Id] INTEGER);"}, expectError: false},
		{input: "Created a Banana!", expected: nil, expectError: true},
		{input: "INSERT INTO Bananas (Id) Values (1),(2),(3);", expected: &Statement{Verb: Insert, SQL: "INSERT INTO Bananas (Id) Values (1),(2),(3);"}, expectError: false},
		{input: "SELECT * FROM Bananas;", expected: nil, expectError: true},
	}

	for _, test := range tests {
		out, err := FromString(test.input)

		if test.expectError && err == nil {
			t.Error("Expected an error but got nothing")
			continue
		}
		if !test.expectError && err != nil {
			t.Errorf("Expected no error but got %s", err)
		}
		if test.expected != nil && (out.Verb != test.expected.Verb || out.SQL != test.expected.SQL) {
			t.Errorf("Expected output %v, but got output %v", test.expected, out)
		}
	}
}

func TestTable(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "CREATE TABLE Creations;", expected: "Creations"},
		{input: "INSERT INTO \"Insertions\" () VALUES ()", expected: "Insertions"},
	}

	for _, test := range tests {
		out, _ := FromString(test.input)
		if out.Table() != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, out.Table())
		}
	}
}
